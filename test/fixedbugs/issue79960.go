// compile


package p

func f[T []byte | []rune]() {
	_ = T("")
}

var _ = f[[]rune]
var _ = f[[]byte]
