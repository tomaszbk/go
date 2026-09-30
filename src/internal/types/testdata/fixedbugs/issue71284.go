package A

type (
	_ = A
	A /* ERROR "invalid recursive type: A refers to itself" */ = A
)
