// run

package main

type SourceRange struct {
	x, y int
}

func (r *SourceRange) String() string {
	return "hello"
}

type SourceNode interface {
	SourceRange()
}

type testNode SourceRange

func (tn testNode) SourceRange() {
}

func main() {
	n := testNode(SourceRange{}) // zero value
	Errorf(n)
}

//go:noinline
func Errorf(n SourceNode) {
	n.SourceRange()
}
