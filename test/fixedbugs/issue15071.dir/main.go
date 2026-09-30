// run


package main

import "os"
import "./exp"

func main() {
	_ = exp.Exported(len(os.Args))
}
