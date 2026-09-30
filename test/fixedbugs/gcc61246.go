// compile

// PR61246: Switch conditions could be untyped, causing an ICE when the
// conditions were lowered into temporaries.
// This is a reduction of a program reported by GoSmith.

package main

func main() {
	switch 1 != 1 {
	default:
	}
}
