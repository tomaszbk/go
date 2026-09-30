package exec

import "io/fs"

// skipStdinCopyError optionally specifies a function which reports
// whether the provided stdin copy error should be ignored.
func skipStdinCopyError(err error) bool {
	// Ignore hungup errors copying to stdin if the program
	// completed successfully otherwise.
	// See Issue 35753.
	pe, ok := err.(*fs.PathError)
	return ok &&
		pe.Op == "write" && pe.Path == "|1" &&
		pe.Err.Error() == "i/o on hungup channel"
}
