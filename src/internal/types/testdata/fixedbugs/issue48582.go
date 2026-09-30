package p

type N /* ERROR "invalid recursive type" */ interface {
	int | N
}

type A /* ERROR "invalid recursive type" */ interface {
	int | B
}

type B interface {
	int | A
}

type S /* ERROR "invalid recursive type" */ struct {
	I // ERROR "interface contains type constraints"
}

type I interface {
	int | S
}

type P interface {
	*P // ERROR "interface contains type constraints"
}
