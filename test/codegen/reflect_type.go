// asmcheck

package codegen

import "reflect"

func intPtrTypeSize() uintptr {
	// amd64:"MOVL [$]8," -"CALL"
	// arm64:"MOVD [$]8," -"CALL"
	return reflect.TypeFor[*int]().Size()
}

func intPtrTypeKind() reflect.Kind {
	// amd64:"MOVL [$]22," -"CALL"
	// arm64:"MOVD [$]22," -"CALL"
	return reflect.TypeFor[*int]().Kind()
}
