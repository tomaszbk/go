// compile


// Issue 22305: gccgo failed to compile this file.

package main

var F func() [0]func()
var i = 2
var B = F()[i]

func main() {}
