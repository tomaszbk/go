// An unexported method can be reachable from the plugin via interface
// when a package is shared. So it need to be live.

package main

import (
	"plugin"

	"testplugin/method3/p"
)

var i p.I

func main() {
	pl, err := plugin.Open("method3.so")
	if err != nil {
		panic(err)
	}

	f, err := pl.Lookup("F")
	if err != nil {
		panic(err)
	}

	f.(func())()

	i = p.T(123)
}
