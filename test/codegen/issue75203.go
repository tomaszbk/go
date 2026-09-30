// asmcheck

package codegen

import "reflect"

func f() reflect.Type {
	// amd64:`LEAQ type:\*int\(SB\)`
	// arm64:`MOVD \$type:\*int\(SB\)`
	return reflect.TypeFor[*int]()
}

func g() reflect.Type {
	// amd64:`LEAQ type:int\(SB\)`
	// arm64:`MOVD \$type:int\(SB\)`
	return reflect.TypeFor[int]()
}
