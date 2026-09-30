// run


package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {
	f()
	panic("deferred function not run")
}

var x = 1

func f() {
	if x == 0 {
		return
	}
	defer g()
	panic("panic")
}

func g() {
	_, file, line, _ := runtime.Caller(2)
	if !strings.HasSuffix(file, "issue5856.go") || line != 28 {
		fmt.Printf("BUG: defer called from %s:%d, want issue5856.go:28\n", file, line)
		os.Exit(1)
	}
	os.Exit(0)
}
