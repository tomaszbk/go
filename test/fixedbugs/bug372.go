// run


// Issue 2355
package main

type T struct {}
func (T) m() string { return "T" }

type TT struct {
	T
	m func() string
}


func ff() string { return "ff" }

func main() {
	var tt TT
	tt.m = ff

	if tt.m() != "ff" {
		println(tt.m(), "!= \"ff\"")
	}
}
