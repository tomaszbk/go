// errorcheck

//go:build 386 || amd64p32 || arm


// Issue 2444

package main
func main() {
	var arr [1000200030]int   // GC_ERROR "type .* too large"
	arr_bkup := arr
	_ = arr_bkup
}
