// errorcheck

// Check that we don't print duplicate errors for string ->
// array-literal conversion

package main

func main() {
	_ = []byte{"foo"}   // ERROR "cannot use|incompatible type|cannot convert"
	_ = []int{"foo"}    // ERROR "cannot use|incompatible type|cannot convert"
	_ = []rune{"foo"}   // ERROR "cannot use|incompatible type|cannot convert"
	_ = []string{"foo"} // OK
}
