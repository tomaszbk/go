package issue52611a

/*
typedef struct Foo {
    int X;
} Foo;
*/
import "C"

func GetX1(foo *C.struct_Foo) int32 {
	return int32(foo.X)
}
