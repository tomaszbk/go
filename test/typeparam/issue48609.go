// compile


package p

func f[T ~chan E, E any](e E) T {
	ch := make(T)
	go func() {
		defer close(ch)
		ch <- e
	}()
	return ch
}
