// run


package main

type Box struct {};
var m map[string] *Box;

func main() {
	m := make(map[string] *Box);
	s := "foo";
	var x *Box = nil;
	m[s] = x;
}

/*
bug058.go:9: illegal types for operand: INDEX
	(MAP[<string>*STRING]*<Box>{})
	(<string>*STRING)
*/
