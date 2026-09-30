// errorcheck

package p

func f(x interface{}) {
	switch x.(type) {
	}

	switch t := x.(type) { // ERROR "declared and not used"
	}
}
