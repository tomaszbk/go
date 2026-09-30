// errorcheck


// Used to crash
// https://golang.org/issue/204

package main

func () x()	// ERROR "no receiver"

func (a b, c d) x()	// ERROR "multiple receiver"

type b int

