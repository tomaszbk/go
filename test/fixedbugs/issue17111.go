// compile


package main

type I int

var (
	i int
	x = I(i)

	e interface{} = x
)
