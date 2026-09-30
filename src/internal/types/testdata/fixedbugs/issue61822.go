// -lang=go1.19

//go:build go1.20

package p

type I[P any] interface {
	~string | ~int
	Error() P
}

func _[P I[string]]() {
	var x P
	var _ error = x
}
