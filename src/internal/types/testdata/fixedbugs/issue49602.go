package p

type M interface {
	m()
}

type C interface {
	comparable
}

type _ interface {
	int | M          // ERROR "cannot use p.M in union (p.M contains methods)"
	int | comparable // ERROR "cannot use comparable in union"
	int | C          // ERROR "cannot use p.C in union (p.C embeds comparable)"
}
