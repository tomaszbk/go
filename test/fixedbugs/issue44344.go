// compile


// Issue #44344: a crash in DWARF scope generation (trying to
// scope the PCs of a function that was inlined away).

package main

func main() {
	pv := []int{3, 4, 5}
	if pv[1] != 9 {
		pv = append(pv, 9)
	}
	tryit := func() bool {
		lpv := len(pv)
		if lpv == 101 {
			return false
		}
		if worst := pv[pv[1]&1]; worst != 101 {
			return true
		}
		return false
	}()
	if tryit {
		println(pv[0])
	}
}
