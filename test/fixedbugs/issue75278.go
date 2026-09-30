// errorcheck -0 -m=2


package p

var a, b []int

func NoIndices() { // ERROR "can inline NoIndices with cost 4 as:.*"
	b = a[:]
}

func LowIndex() { // ERROR "can inline LowIndex with cost 4 as:.*"
	b = a[0:]
}

func HighIndex() { // ERROR "can inline HighIndex with cost 4 as:.*"
	b = a[:len(a)]
}

func BothIndices() { // ERROR "can inline BothIndices with cost 4 as:.*"
	b = a[0:len(a)]
}
