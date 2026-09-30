// errorcheck


package foo

var i int

func (*i) bar() // ERROR "not a type|expected type"

