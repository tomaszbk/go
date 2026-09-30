// errorcheck


package p

type T interface { // ERROR "invalid recursive type: anonymous interface refers to itself"
	M(interface {
		T
	})
}
