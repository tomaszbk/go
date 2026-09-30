package p

import (
	"fmt"
)

func f() {
	int status // ERROR syntax error: unexpected name status at end of statement
	fmt.Println(status)
}
