// compile

// PR61253: gccgo incorrectly parsed the
// `RecvStmt = ExpressionList "=" RecvExpr` production.

package main

func main() {
	c := make(chan int)
	v := new(int)
	b := new(bool)
	select {
	case (*v), (*b) = <-c:
	}

}
