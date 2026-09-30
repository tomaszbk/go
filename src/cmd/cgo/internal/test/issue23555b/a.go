package issue23555

// #include <stdlib.h>
import "C"

func X() {
	C.free(C.malloc(10))
}
