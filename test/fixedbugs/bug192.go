// errorcheck


package main

import "fmt"	// GCCGO_ERROR "previous"

var _ = fmt.Println // avoid imported and not used error

var fmt int	// ERROR "redecl|redefinition|fmt already declared"
