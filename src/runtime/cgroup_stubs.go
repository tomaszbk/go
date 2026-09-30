//go:build !linux

package runtime

func defaultGOMAXPROCSInit()          {}
func defaultGOMAXPROCSUpdateGODEBUG() {}

func defaultGOMAXPROCS(ncpu int32) int32 {
	// Use the total number of logical CPUs available now, as CPU affinity
	// may change after start.
	//
	// TODO(prattmic): On some GOOS getCPUCount can never change. Don't
	// bother calling over and over.

	procs := ncpu
	if procs <= 0 {
		procs = getCPUCount()
	}
	return procs
}
