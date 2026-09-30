// compile


package p

func hashBytesRaw(b0, b1, b2, b3, b7 byte) uint64 {
	return (uint64(b0) | uint64(b1)<<8 | uint64(b2)<<16 | uint64(b3)<<24)
}

func doStuff(data []byte) uint64 {
	return hashBytesRaw(data[0], data[1], data[2], data[3], data[7])
}
