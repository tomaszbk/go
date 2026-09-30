package p

func _[P struct{ f int }](x P) {
	_ = x.g // ERROR "type P has no field or method g"
}

func _[P struct{ f int } | struct{ g int }](x P) {
	_ = x.g // ERROR "type P has no field or method g"
}
