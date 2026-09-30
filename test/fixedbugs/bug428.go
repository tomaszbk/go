// run

// Test that when the compiler expands append inline it does not
// overwrite a value before it needs it (issue 3369).

package main

func main() {
	s := make([]byte, 5, 6)
	copy(s, "12346")
	s = append(s[:len(s)-1], '5', s[len(s)-1])
	if string(s) != "123456" {
		panic(s)
	}
}
