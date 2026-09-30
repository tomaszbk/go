package main

import "./bug"

type foo int

func (f *foo) Bar() {
}

func main() {
	bug.Foo(new(foo))
}
