// errorcheck


// https://golang.org/issue/808

package main

type A [...]int	// ERROR "outside of array literal|invalid use of \[\.\.\.\]"


