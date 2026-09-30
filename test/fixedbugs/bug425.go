// compile


// https://golang.org/issue/3119

package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println(s == "")
	fmt.Println(s + "world" == "world")
}
