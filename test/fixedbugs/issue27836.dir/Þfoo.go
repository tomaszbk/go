package Þfoo

var ÞbarV int = 101

func Þbar(x int) int {
	defer func() { ÞbarV += 3 }()
	return Þblix(x)
}

func Þblix(x int) int {
	defer func() { ÞbarV += 9 }()
	return ÞbarV + x
}
