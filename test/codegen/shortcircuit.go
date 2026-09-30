// asmcheck

package codegen

func efaceExtract(e interface{}) int {
	// This should be compiled with only
	// a single conditional jump.
	// amd64:-"JMP"
	if x, ok := e.(int); ok {
		return x
	}
	return 0
}
