// compile

// issue 11370: cmd/compile: "0"[0] should not be a constant

package p

func main() {
	println(-"abc"[1] >> 1)
}
