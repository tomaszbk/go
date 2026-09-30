// compile


// The gccgo compiler used to crash building a comparison between an
// interface and an empty struct literal.

package p
 
type S struct{}

func F(v interface{}) bool {
	return v == S{}
}
