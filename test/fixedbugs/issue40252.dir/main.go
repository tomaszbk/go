package main

import "./a"

func main() {
	defer func() {
		if recover() == nil {
			panic("expected nil pointer dereference")
		}
	}()
	a.Call()
}
