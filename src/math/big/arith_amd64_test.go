//go:build !math_big_pure_go

package big

import "testing"

func TestAddMulVVWWNoADX(t *testing.T) {
	setDuringTest(t, &hasADX, false)
	TestAddMulVVWW(t)
}
