// errorcheck

package p

import _ "embed"

func f() {
	//go:embed x.txt // ERROR "go:embed cannot apply to var inside func"
	var x string
	_ = x
}
