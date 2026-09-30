package repro

import "testing"

func TestSomething(t *testing.T) {
	small()
	for i := 0; i < 1001; i++ {
		large(i)
	}
}
