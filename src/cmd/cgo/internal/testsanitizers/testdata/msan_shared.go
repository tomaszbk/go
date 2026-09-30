// This program segfaulted during libpreinit when built with -msan:
// http://golang.org/issue/18707

package main

import "C"

func main() {}
