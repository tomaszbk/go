package issue29563

//int foo1();
//int foo2();
import "C"

func Bar() int {
	return int(C.foo1()) + int(C.foo2())
}
