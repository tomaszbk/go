// run

package main

var ok [2]bool

func main() {
	f()()
	if !ok[0] || !ok[1] {
		panic("FAIL")
	}
}

func f() func() { ok[0] = true; return g }
func g()        { ok[1] = true }
