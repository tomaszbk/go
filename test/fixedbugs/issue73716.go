// build

// Issue 73716: cmd/compile: unnamed functions missing FuncInfo

package main

import "fmt"

type EP func()
type F func(EP) EP

func main() {
	eps := []EP{ep1, ep2}
	var h EP

	for _, ep := range eps {
		h = F(func(e EP) EP {
			return func() {
				ep()
				e()
			}
		})(h)
	}
	h()
}

func ep1() {
	fmt.Printf("ep1\n")
}

func ep2() {
	fmt.Printf("ep2\n")
}
