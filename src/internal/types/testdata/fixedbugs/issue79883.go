package p

func _[P []byte](p string) {
	_ = (*P)(p /* ERROR "pointer to type parameter" */)
}
