package main

import "featuretest/lib"

func main() {
	if lib.Pick(2)(3) != 5 || lib.Generic("ok")() != "ok" {
		panic("export closures")
	}
}
