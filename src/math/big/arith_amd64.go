//go:build !math_big_pure_go

package big

import "internal/cpu"

var hasADX = cpu.X86.HasADX && cpu.X86.HasBMI2
