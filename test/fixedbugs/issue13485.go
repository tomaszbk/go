// errorcheck


package p

var (
	_ [10]int
	_ [10.0]int
	_ [float64(10)]int                // ERROR "invalid array bound|must be integer"
	_ [10 + 0i]int
	_ [complex(10, 0)]int
	_ [complex128(complex(10, 0))]int // ERROR "invalid array bound|must be integer"
	_ ['a']int
	_ [rune(65)]int
)
