// run


// Gccgo mishandles converting an untyped boolean to an interface type.

package main

func t(args ...interface{}) bool {
        x := true
        return x == args[0]
}

func main() {
	r := t("x" == "x" && "y" == "y")
	if !r {
		panic(r)
	}
}
