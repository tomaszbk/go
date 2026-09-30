package issue52611a

import "C"

func GetX2(foo *C.struct_Foo) int32 {
	return int32(foo.X)
}
