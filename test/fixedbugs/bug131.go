// errorcheck

package main

func main() {
	const a uint64 = 10
	var _ int64 = a // ERROR "convert|cannot|incompatible"
}
