// Issue 62430: a program that uses plugins may appear
// to have no references to an initialized global map variable defined
// in some stdlib package (ex: unicode), however there
// may be references to that map var from a plugin that
// gets loaded.

package main

import (
	"fmt"
	"plugin"
	"unicode"
)

func main() {
	p, err := plugin.Open("issue62430.so")
	if err != nil {
		panic(err)
	}
	s, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}

	f := s.(func(string) *unicode.RangeTable)
	if f("C") == nil {
		panic("unicode.Categories not properly initialized")
	} else {
		fmt.Println("unicode.Categories properly initialized")
	}
}
