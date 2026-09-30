// compile

package main

import "runtime"

func f(x interface{}) {
	runtime.KeepAlive(x)
}
