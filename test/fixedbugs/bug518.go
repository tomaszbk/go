// errorcheck


// The gofrontend used to accept this.

package p

func F2(a int32) bool {
	return a == C	// ERROR "invalid|incompatible"
}

const C = uint32(34)
