// errorcheck


package p

func _() {
	var ch chan bool
	select {
	default:
	case <-ch { // GCCGO_ERROR "expected colon"
	}           // GC_ERROR "expected :"
}
