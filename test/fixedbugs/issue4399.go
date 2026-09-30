// compile


// Issue 4399: 8g would print "gins LEAQ nil *A".

package main

type A struct{ a int }

func main() {
	println(((*A)(nil)).a)
}
