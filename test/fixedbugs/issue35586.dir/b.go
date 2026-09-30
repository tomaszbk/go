package b

import "./a"

func F(addr string) (uint64, string) {
	return a.D(addr, 32)
}
