package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
)

//go:embed config.json
var configFile string

//go:embed api.json.tmpl
var templateFile string

func main() {

	var config config
	if err := json.Unmarshal([]byte(configFile), &config); err != nil {
		die(fmt.Errorf("cannot unmarshal config file %q: %s", preview(configFile), err))
	}

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
	Paths map[string]struct {
		Create bool `json:"create"`
		Update bool `json:"update"`
		List   bool `json:"list"`
		Get    bool `json:"get"`
		Watch  bool `json:"watch"`
		Delete bool `json:"delete"`
	} `json:"paths"`
}
