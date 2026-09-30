package main

// int the_answer();
import "C"
import (
	"fmt"
	"os"
)

func TheAnswer() int {
	return int(C.the_answer())
}

func main() {
	if a := TheAnswer(); a != 42 {
		fmt.Fprintln(os.Stderr, "Unexpected result for The Answer. Got:", a, " Want: 42")
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "ok")
}
