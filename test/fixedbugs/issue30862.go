// runindir -goexperiment fieldtrack

// Test case for issue 30862.  This test as written will
// fail for the main 'gc' compiler unless GOEXPERIMENT=fieldtrack
// is set when building it, whereas gccgo has field tracking
// enabled by default (hence the build tag below).

package ignored
