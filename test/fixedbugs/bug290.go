// run


// https://golang.org/issue/920

package main

type X struct { x []X }

func main() {
	type Y struct { x []Y }	// used to get invalid recursive type
}
