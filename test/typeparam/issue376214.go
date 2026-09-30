// run


package main

func add[S ~string | ~[]byte](buf *[]byte, s S) {
	*buf = append(*buf, s...)
}

func main() {
	var buf []byte
	add(&buf, "foo")
	add(&buf, []byte("bar"))
	if string(buf) != "foobar" {
		panic("got " + string(buf))
	}
}
