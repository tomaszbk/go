//go:build !math_big_pure_go

package big

import "testing"

func TestAddVVNoVec(t *testing.T) {
	setDuringTest(t, &hasVX, false)
	TestAddVV(t)
}

func TestSubVVNoVec(t *testing.T) {
	setDuringTest(t, &hasVX, false)
	TestSubVV(t)
}
