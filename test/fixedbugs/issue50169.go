// compile

package main

func main() {
	var x Value
	NewScanner().Scan(x)
}

type Value any

type Scanner interface{ Scan(any) error }

func NewScanner() Scanner {
	return &t{}
}

type t struct{}

func (*t) Scan(interface{}) error { return nil }
