package p

type (
	_ [1 /* ERROR "invalid array length" */ << 100]int
	_ [1.0]int
	_ [1.1 /* ERROR "must be integer" */ ]int
)
