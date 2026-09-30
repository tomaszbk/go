// asmcheck

package codegen

func a(n string) bool {
	// arm64:"CBZ"
	if len(n) > 0 {
		return true
	}
	return false
}

func a2(n []int) bool {
	// arm64:"CBZ"
	if len(n) > 0 {
		return true
	}
	return false
}

func a3(n []int) bool {
	// amd64:"TESTQ"
	if len(n) < 1 {
		return true
	}
	return false
}
