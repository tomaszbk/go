// compile

//go:build amd64

package p

var _ = []int{1 << 31: 1} // ok on machines with 64bit int
