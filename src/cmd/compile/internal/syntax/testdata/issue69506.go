package p

func _(a int, b /* ERROR missing parameter type */ )
func _(a int, /* ERROR missing parameter name */ []int)
func _(a int, /* ERROR missing parameter name */ []int, c int)
