// errorcheck


package main

func main() {
	type Slice []byte;
	a := [...]byte{ 0 };
	b := Slice(a[0:]);	// This should be OK.
	c := Slice(a);		// ERROR "invalid|illegal|cannot"
	_, _ = b, c;
}
