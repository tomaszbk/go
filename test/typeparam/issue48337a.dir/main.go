package main

import "./a"

func main() {
	obj := a.NewWrapperWithLock("this file does import sync")
	obj.PrintWithLock()
}
