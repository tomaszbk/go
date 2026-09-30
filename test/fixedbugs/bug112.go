// compile


package main

type T struct { s string }
var t = T{"hi"}

func main() {}

/*
bug112.go:6: illegal conversion of constant to T
*/
