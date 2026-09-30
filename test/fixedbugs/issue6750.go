// errorcheck


package main

import "fmt"

func printmany(nums ...int) {
	for i, n := range nums {
		fmt.Printf("%d: %d\n", i, n)
	}
	fmt.Printf("\n")
}

func main() {
	printmany(1, 2, 3)
	printmany([]int{1, 2, 3}...)
	printmany(1, "abc", []int{2, 3}...) // ERROR "too many arguments in call( to printmany\n\thave \(number, string, \.\.\.int\)\n\twant \(...int\))?"
}
