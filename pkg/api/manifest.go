// Package api provides the spirits API. These types are used to interact with the spirits command
// line and game server.
//
// See runtime packages for more information about each type (e.g., pkg/manifest for Manifest,
// pkg/team for Team, pkg/spirit for Spirit, pkg/action for Action, etc.).
package api

type Manifest struct {
	Metadata map[string]string `json:"metadata"`
	Data     *ManifestData     `json:"data"`
}

type ManifestData struct {
	Teams       []*Team       `json:"teams"`
	Validations []*Validation `json:"validations"`
}

type Team struct {
	Name    string    `json:"name"`
	Spirits []*Spirit `json:"spirits"`
}

type Spirit struct {
	Name string `json:"name"`

	Health  int `json:"health"`
	Power   int `json:"power"`
	Armour  int `json:"armour"`
	Agility int `json:"agility"`

	Action *Action `json:"action"`
}

type Action struct {
	Type   string `json:"type"`
	Target int    `json:"target"`
}

type Validation struct {
	Name string `json:"name"`
}
