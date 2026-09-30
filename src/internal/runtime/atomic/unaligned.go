package atomic

func panicUnaligned() {
	panic("unaligned 64-bit atomic operation")
}
