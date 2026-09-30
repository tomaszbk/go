// compile


// PR61264: IncDec statements involving composite literals caused in ICE in gccgo.

package main

func main() {
        map[int]int{}[0]++
}
