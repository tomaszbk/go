//go:build !math_big_pure_go

package big

import "internal/cpu"

var hasVX = cpu.S390X.HasVX

func addVVvec(z, x, y []Word) (c Word)
func subVVvec(z, x, y []Word) (c Word)
