// compile

// Gccgo used to give an incorrect error
// bug495.go:16:2: error: missing statement after label

package p

func F(i int) {
	switch i {
	case 0:
		goto lab
	lab:
		fallthrough
	case 1:
	}
}
