// errorcheck


// Issue 9634: Structs are incorrectly unpacked when passed as an argument
// to append.

package main

func main() {
	s := struct{
		t []int
		u int
	}{}
	_ = append(s, 0) // ERROR "must be a slice|must be slice|not a slice"
}
