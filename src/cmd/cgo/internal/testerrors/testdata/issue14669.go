// Issue 14669: test that fails when build with CGO_CFLAGS selecting
// optimization.

package p

/*
const int E = 1;

typedef struct s {
	int       c;
} s;
*/
import "C"

func F() {
	_ = C.s{
		c: C.E,
	}
}
