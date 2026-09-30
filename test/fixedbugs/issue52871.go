// compile


// gofrontend crashed compiling this code.

package p

type S struct {}

func (s *S) test(_ string) {}

var T = [1]func(*S, string) {
	(*S).test,
}
