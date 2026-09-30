package main

import p "./b"

var G int

func main() {
	if G == 101 {
		p.G(nil, nil)
	}
}
