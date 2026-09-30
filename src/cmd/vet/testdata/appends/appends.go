// This file contains tests for the appends checker.

package appends

func AppendsTest() {
	sli := []string{"a", "b", "c"}
	sli = append(sli) // ERROR "append with no values"
}
