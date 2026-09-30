// Package startlinetest contains helpers for runtime_test.TestStartLineAsm.
package startlinetest

// Defined in func_amd64.s, this is a trivial assembly function that calls
// runtime_test.callerStartLine.
func AsmFunc() int

// Provided by runtime_test.
var CallerStartLine func(bool) int
