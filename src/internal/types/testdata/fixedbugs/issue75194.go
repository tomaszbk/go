package p

type A /* ERROR "invalid recursive type: A refers to itself" */ struct {
	a A
}

type B /* ERROR "invalid recursive type: B refers to itself" */ struct {
	a A
	b B
}
