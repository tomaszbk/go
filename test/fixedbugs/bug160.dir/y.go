package main

import "os"
import "./x"

func main() {
	if x.Zero != 0 {
		println("x.Zero = ", x.Zero);
		os.Exit(1);
	}
	if x.Ten != 10 {
		println("x.Ten = ", x.Ten);
		os.Exit(1);
	}
}
