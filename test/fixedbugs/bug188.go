// errorcheck -d=panic

package main

import "sort"

func main() {
	sort.Sort(nil)
	var x int
	sort(x) // ERROR "package"
}
