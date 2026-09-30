// compile

// Compile with static map literal.

package p

type i interface {
	j()
}

type s struct{}

func (s) j() {}

type foo map[string]i

var f = foo{
	"1": s{},
	"2": s{},
}
