//go:build !race

package quic

func raceAcquire()      {}
func raceReleaseMerge() {}
