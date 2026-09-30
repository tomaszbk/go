package race_test

// golang.org/issue/13264
// The test is that this compiles at all.

func issue13264() {
	for ; ; []map[int]int{}[0][0] = 0 {
	}
}
