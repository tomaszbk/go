package p

func _[S string | []byte](s S) {
	var buf []byte
	_ = append(buf, s...)
}

func _[S ~string | ~[]byte](s S) {
	var buf []byte
	_ = append(buf, s...)
}

// test case from issue

type byteseq interface {
	string | []byte
}

// This should allow to eliminate the two functions above.
func AppendByteString[source byteseq](buf []byte, s source) []byte {
	return append(buf, s[1:6]...)
}
