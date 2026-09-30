package p

func _() {
	len.Println /* ERROR "invalid use of len (built-in) in selector expression" */
	len.Println /* ERROR "invalid use of len (built-in) in selector expression" */ ()
	_ = len.Println /* ERROR "invalid use of len (built-in) in selector expression" */
	_ = len /* ERROR "cannot index len" */ [0]
	_ = *len /* ERROR "cannot indirect len" */
}
