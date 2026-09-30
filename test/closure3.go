// errorcheckandrundir -0 -m -d=inlfuncswithclosures=1

//go:build !goexperiment.newinliner

// Check correctness of various closure corner cases
// that are expected to be inlined

package ignored
