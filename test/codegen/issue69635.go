// asmcheck


package codegen

func calc(a uint64) uint64 {
	v := a >> 20 & 0x7f
	// amd64: `SHRQ \$17, AX$` `ANDL \$1016, AX$`
	return v << 3
}
