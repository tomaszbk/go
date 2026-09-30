package main

import "./a"

func F() any {
	return struct{ int }{0}
}

func main() {
	_, ok1 := F().(struct{ int })
	_, ok2 := a.F().(struct{ int })
	if !ok1 || ok2 {
		panic(0)
	}
}
