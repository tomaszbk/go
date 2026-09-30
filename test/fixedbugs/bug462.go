// errorcheck

package main

import "os"

var _ = os.Open // avoid imported and not used error

type T struct {
	File int
}

func main() {
	_ = T{
		os.File: 1, // ERROR "invalid field name os.File|unknown field"
	}
}
