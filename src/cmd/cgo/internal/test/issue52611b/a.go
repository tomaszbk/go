package issue52611b

import "C"

func GetX1(bar *C.struct_Bar) int32 {
	return int32(bar.X)
}
