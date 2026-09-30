// run

// Test defer.

package main

import "fmt"

var result string

func addInt(i int) { result += fmt.Sprint(i) }

func test1helper() {
	for i := 0; i < 10; i++ {
		defer addInt(i)
	}
}

func test1() {
	result = ""
	test1helper()
	if result != "9876543210" {
		fmt.Printf("test1: bad defer result (should be 9876543210): %q\n", result)
		panic("defer")
	}
}

func addDotDotDot(v ...interface{}) { result += fmt.Sprint(v...) }

func test2helper() {
	for i := 0; i < 10; i++ {
		defer addDotDotDot(i)
	}
}

func test2() {
	result = ""
	test2helper()
	if result != "9876543210" {
		fmt.Printf("test2: bad defer result (should be 9876543210): %q\n", result)
		panic("defer")
	}
}

func main() {
	test1()
	test2()
}
