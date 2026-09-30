// A type can be passed to a plugin and converted to interface
// there. So its methods need to be live.

package main

import (
	"plugin"

	"testplugin/method2/p"
)

var t p.T

type I interface{ M() }

func main() {
	pl, err := plugin.Open("method2.so")
	if err != nil {
		panic(err)
	}

	f, err := pl.Lookup("F")
	if err != nil {
		panic(err)
	}

	f.(func(p.T) interface{})(t).(I).M()
}
