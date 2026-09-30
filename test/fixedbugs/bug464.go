// errorcheck


// Issue 3937: unhelpful typechecking loop message
// for identifiers wrongly used as types.

package main

func foo(x foo) {} // ERROR "expected type|not a type"
