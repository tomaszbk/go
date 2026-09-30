package flag

import (
	"io"
	"os"
)

// Additional routines compiled into the package only during testing.

var DefaultUsage = Usage

// ResetForTesting clears all flag state and sets the usage function as directed.
// After calling ResetForTesting, parse errors in flag handling will not
// exit the program.
func ResetForTesting(usage func()) {
	CommandLine = NewFlagSet(os.Args[0], ContinueOnError)
	CommandLine.SetOutput(io.Discard)
	CommandLine.Usage = commandLineUsage
	Usage = usage
}
