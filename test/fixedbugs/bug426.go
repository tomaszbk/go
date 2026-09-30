// compile


// gccgo crashed compiling this.

package p

type T *T

func f(t T) {
	println(t, *t)
}
