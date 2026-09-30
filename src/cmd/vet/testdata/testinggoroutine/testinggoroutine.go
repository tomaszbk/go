package testinggoroutine

import "testing"

func _(t *testing.T) {
	go func() {
		t.Fatal("fail") // ERROR `call to \(\*testing.T\).Fatal from a non-test goroutine`
	}()
}
