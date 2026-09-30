// compile -goexperiment fieldtrack

package p

func a(x struct{ f int }) { _ = x.f }

func b() { a(struct{ f int }{}) }
