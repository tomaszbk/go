// compile

package p

type big struct {
	_ [1 << 24]byte
	f [2]float32
}

var x big

func load() [2]float32   { return x.f }
func store(v [2]float32) { x.f = v }
