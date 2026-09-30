package main

/*
const int sizeofLongDouble = sizeof(long double);
*/
import "C"

import "fmt"

func main() {
	fmt.Println(C.sizeofLongDouble)
}
