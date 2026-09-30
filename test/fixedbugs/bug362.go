// errorcheck


// issue 1662
// iota inside var

package main

var (
	a = iota  // ERROR "undefined: iota|iota is only defined in const|cannot use iota outside constant declaration"
	b = iota  // ERROR "undefined: iota|iota is only defined in const|cannot use iota outside constant declaration"
	c = iota  // ERROR "undefined: iota|iota is only defined in const|cannot use iota outside constant declaration"
)
