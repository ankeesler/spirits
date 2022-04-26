package version

import _ "embed"

// TODO: this should come from api.JSON so it is only defined a single place
//go:embed version.txt
var Version string
