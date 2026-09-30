// errorcheck


package main

func f(x int, y ...int) // ok

func g(x int, y float32) (...)	// ERROR "[.][.][.]"

var x ...int;		// ERROR "[.][.][.]|syntax|type"

type T ...int;		// ERROR "[.][.][.]|syntax|type"
