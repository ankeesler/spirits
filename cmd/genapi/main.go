package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	pathpkg "path"
	"strings"
)

//go:embed config.json
var configFile string

//go:embed api.json.tmpl
var templateFile string

func main() {
	if len(os.Args) != 2 {
		die(fmt.Errorf("usage: %s <version>", os.Args[0]))
	}
	version := os.Args[1]

	var config config
	if err := json.Unmarshal([]byte(configFile), &config); err != nil {
		die(fmt.Errorf("cannot unmarshal config file %q: %s", preview(configFile), err))
	}
	config.Version = version

	tmpl, err := template.New("api.json.tmpl").Parse(templateFile)
	if err != nil {
		die(fmt.Errorf("cannot parse template file %q: %s", preview(templateFile), err))
	}

	if err := tmpl.Execute(os.Stdout, &config); err != nil {
		die(fmt.Errorf("cannot execute template %q: %s", tmpl.Name(), err))
	}
}

func die(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func preview(s string) string {
	if len(s) < 12 {
		return s
	}
	return s[:6] + "..." + s[len(s)-6:]
}

type config struct {
	Version string

	Paths []path `json:"paths"`
}

type path struct {
	Name  string   `json:"name"`
	Verbs []string `json:"verbs"`
}

func (p *path) Plural() string {
	base := pathpkg.Base(p.Name)
	if isParam(base) {
		var prefix string
		prefix, base = pathpkg.Split(p.Name)
		base = pathpkg.Base(prefix)
	}
	return strings.Title(base)
}

func (p *path) Singular() string {
	plural := p.Plural()
	return plural[:len(plural)-1]
}

func (p *path) Params() []string {
	var params []string
	for _, segment := range strings.Split(p.Name, "/") {
		if isParam(segment) {
			params = append(params, extractParam(segment))
		}
	}
	return params
}

func (p *path) Scopes(read bool) []string {
	baseScope := "spirits:" + strings.Join(p.resources(), ":")

	var scopes []string
	scopes = append(scopes, baseScope+".write")
	if read {
		scopes = append(scopes, baseScope+".read")
	}

	return scopes
}

func (p *path) OperationIDSuffix() string {
	suffix := ""
	resources := p.resources()
	for i, resource := range resources {
		suffix += strings.Title(resource)

		// Every word but the last should be single
		if len(suffix) > 0 && i != len(resources)-1 {
			suffix = suffix[:len(suffix)-1]
		}
	}
	return suffix
}

func (p *path) resources() []string {
	var tags []string

	for _, segment := range strings.Split(p.Name, "/") {
		if len(segment) > 0 && !isParam(segment) {
			tags = append(tags, segment)
		}
	}

	return tags
}

func isParam(s string) bool {
	if len(s) == 0 {
		return false
	}
	return s[0] == '{' && s[len(s)-1] == '}'
}

func extractParam(s string) string {
	if len(s) <= 1 {
		return ""
	}
	return s[1 : len(s)-1]
}
