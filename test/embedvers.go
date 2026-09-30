// errorcheck -lang=go1.15

package p

import _ "embed"

//go:embed x.txt // ERROR "go:embed requires go1.16 or later"
var x string
