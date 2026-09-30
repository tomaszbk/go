// compile

// This code relies on pre-1.28 string(integer) conversion rules.
//go:build !go1.28

package p

const _ = -uint(len(string(1<<32)) - len("\uFFFD"))
