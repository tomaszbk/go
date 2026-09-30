// errorcheckandrundir -1

//go:build !nacl && !js && !plan9

package ignored

// Compile: bug0.go, bug1.go
// Compile and errorCheck: bug2.go
// Link and run: bug3.go
