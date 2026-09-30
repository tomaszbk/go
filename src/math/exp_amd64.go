//go:build amd64

package math

import "internal/cpu"

var useFMA = cpu.X86.HasAVX && cpu.X86.HasFMA
