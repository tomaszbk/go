//go:build netgo && netcgo

package net

func init() {
	// This will give a compile time error about the unused constant.
	// The advantage of this approach is that the gc compiler
	// actually prints the constant, making the problem obvious.
	"Do not use both netgo and netcgo build tags."
}
