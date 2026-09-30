// compile

package a

type I[T any] interface{ M() T }

var _ = I[int].M
