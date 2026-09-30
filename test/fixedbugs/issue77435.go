// compile


// Issue 77435: compiler crash on clear of map resulting
// from a map lookup (or some other syntax that is
// non-idempotent during walk).

package p

func f(s map[int]map[int]int) {
	clear(s[0])
}
