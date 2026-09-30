// run

package main

func main() {
	for i := 0; i < 10000; i++ {
		h(i)
		sink = make([]byte, 1024) // generate some garbage
	}
}

func h(iter int) {
	var x [32]byte
	for i := 0; i < 32; i++ {
		x[i] = 99
	}
	g(&x)
	if x == ([32]byte{}) {
		return
	}
	for i := 0; i < 32; i++ {
		println(x[i])
	}
	panic(iter)
}

//go:noinline
func g(x interface{}) {
	switch e := x.(type) {
	case *[32]byte:
		var c [32]byte
		*e = c
	case *[]byte:
		*e = nil
	}
}

var sink []byte
