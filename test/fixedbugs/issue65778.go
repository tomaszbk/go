// compile -godebug gotypesalias=1


package p

type A = int

type T[P any] *A

var _ T[int]
