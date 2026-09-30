package p

var badSlice [8265]byte

func init() {
	badSlice[0] = 4
}
