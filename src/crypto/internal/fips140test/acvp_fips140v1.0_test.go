//go:build fips140v1.0

package fipstest

import _ "embed"

//go:embed acvp_capabilities_fips140v1.0.json
var capabilitiesJson []byte

var testConfigFile = "acvp_test_fips140v1.0.config.json"
