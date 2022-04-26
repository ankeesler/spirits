package api

import (
	_ "embed"
	"encoding/json"
)

//go:embed generated/api.json
var JSON string

var Version string

func init() {
	var openapi struct {
		Info struct {
			Version string `json:"version"`
		} `json:"info"`
	}
	if err := json.Unmarshal([]byte(JSON), &openapi); err != nil {
		panic(err)
	}
	Version = openapi.Info.Version
}
