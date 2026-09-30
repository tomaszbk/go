package issue76861

// #cgo CFLAGS: -Wall -Wextra -Werror
// void issue76861(void) {}
import "C"

func Issue76861() {
	C.issue76861()
}
