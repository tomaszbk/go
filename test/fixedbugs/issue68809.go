// run

package main

func main() {
	cnt := 0
	for i := 1; i <= 11; i++ {
		if i-6 > 4 {
			cnt++
		}
	}
	if cnt != 1 {
		panic("bad")
	}
}
