package p

import "strings"

type T struct{}

func (T) M() {
	strings.HasPrefix("", "")
}
