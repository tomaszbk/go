// errorcheck


// Checking that line number is correct in error message.

package main

type Cint int

func foobar(*Cint, Cint, Cint, *Cint)

func main() {
	a := Cint(1)

	foobar(
		&a,
		0,
		0,
		42, // ERROR ".*"
	)
}
