// run

// This code relies on pre-1.28 string(integer) conversion rules.
//go:build !go1.28


package main

func main() {
	const fffd = "\uFFFD"

	// runtime.intstring used to convert int64 to rune without checking
	// for truncation.
	u := uint64(0x10001f4a9)
	big := string(u)
	if big != fffd {
		panic("big != bad")
	}

	// cmd/compile used to require integer constants to fit into an "int".
	const huge = string(1 << 100)
	if huge != fffd {
		panic("huge != bad")
	}
}
