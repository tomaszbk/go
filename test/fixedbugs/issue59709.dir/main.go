package main

import (
	"./dcache"
)

func main() {
	var m dcache.Module
	m.Configure("x")
	m.Configure("y")
	var e error
	m.Blurb("x", e)
}
