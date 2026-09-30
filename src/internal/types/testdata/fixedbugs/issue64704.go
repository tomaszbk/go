// -lang=go1.21


package p

func _() {
	for range 10 /* ERROR "cannot range over 10 (untyped int constant): requires go1.22 or later" */ {
	}
}
