// errorcheck -0 -m=2

package main

//go:noinline
func run() { // ERROR "cannot inline run: marked go:noinline"
	f := func() { // ERROR "can inline run.func1 with cost .* as:.*" "func literal does not escape"
		g() // ERROR "inlining call to g"
	}
	f() // ERROR "inlining call to run.func1" "inlining call to g"
	_ = f
	run()
}

func g() { // ERROR "can inline g with cost .* as:.*"
}

func main() { // ERROR "can inline main with cost .* as:.*"
	run()
}
