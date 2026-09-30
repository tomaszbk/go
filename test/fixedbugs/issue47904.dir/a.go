package a

type K [2]int

// Package a only builds the slice literal, so the [2]K array walk builds
// for it is noalg here, and a's descriptor for [2]K has no algorithms.
var Sink any

func init() { Sink = []K{{1, 2}, {3, 4}} }
