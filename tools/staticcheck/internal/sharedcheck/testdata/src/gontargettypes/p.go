package gontargettypes

type node struct{ P *int }

func f(flag bool, p *int, n *node) {
	var callback func(int) int = (x) => x + 1
	var conditional any = if flag { p } else { nil }
	var chain any = n?.P
	var fallback any = p ?? new(int)
	var legacy int = 1 // want "should omit type int"
	_, _, _, _, _ = callback, conditional, chain, fallback, legacy
}
