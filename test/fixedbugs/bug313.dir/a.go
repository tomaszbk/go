package main

import "fmt"

func a() {
	fmt.DoesNotExist() // ERROR "undefined"
}
