package api

import (
	_ "embed"
	"encoding/json"
)

//go:embed generated/api.json
var JSON string

var Object map[string]interface{}

var Version string

func init() {
	if err := json.Unmarshal([]byte(JSON), &Object); err != nil {
		panic(err)
	}
	Version = Object["info"].(map[string]interface{})["version"].(string)
}
