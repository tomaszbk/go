package main

import "./a"

func main() {
	_ = p.S{1} // ERROR "too few values.*unexported fields"
}
