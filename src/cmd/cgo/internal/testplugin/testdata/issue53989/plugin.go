package main

import "testplugin/issue53989/p"

func Square(x int) { // export Square for plugin
	p.Square(x)
}

func main() {}
