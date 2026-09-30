package a

import "a/imported"

func _() {
	var x int
	imported.F[int](x) // want "unnecessary type arguments"
}
