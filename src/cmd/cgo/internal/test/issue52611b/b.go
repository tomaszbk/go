package issue52611b

/*
typedef struct Bar {
    int X;
} Bar;
*/
import "C"

func GetX2(bar *C.struct_Bar) int32 {
	return int32(bar.X)
}
