// run

package main

func main() {
	f[interface{}](nil)
}

func f[T any](x T) {
	var _ interface{} = x
}
