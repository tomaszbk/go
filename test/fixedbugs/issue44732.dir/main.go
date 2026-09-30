package main

import (
	"issue44732.dir/bar"
	"issue44732.dir/foo"
)

func main() {
	_ = bar.Bar{}
	_ = foo.NewFoo()
}
