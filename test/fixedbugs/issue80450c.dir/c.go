package main

import (
	target "./a"
	workflow "./b"
	_ "reflect"
)

var _ = workflow.Const(target.P)

func main() {}
