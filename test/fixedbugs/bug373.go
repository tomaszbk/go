// errorcheck


// Issue 873, 2162

package foo

func f(x interface{}) {
	switch t := x.(type) {  // ERROR "declared and not used"
	case int:
	}
}

func g(x interface{}) {
	switch t := x.(type) {
	case int:
	case float32:
		println(t)
	}
}

func h(x interface{}) {
	switch t := x.(type) {
	case int:
	case float32:
	default:
		println(t)
	}
}
