// run

package main

const maxUint64 = (1 << 64) - 1

//go:noinline
func f(n uint64) uint64 {
	return maxUint64 - maxUint64%n
}

func main() {
	for i := uint64(1); i < 20; i++ {
		println(i, maxUint64-f(i))
	}
}
