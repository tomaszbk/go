// compile


// issue 1908
// unreasonable width used to be internal fatal error

package test

func main() {
	buf := [1<<30]byte{}
	_ = buf[:]
}
