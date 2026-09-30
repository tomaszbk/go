// compile


package p

func f(b []byte) []byte {
	return (*[32]byte)(b[:32])[:]
}
