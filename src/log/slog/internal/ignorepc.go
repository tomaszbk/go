package internal

// If IgnorePC is true, do not invoke runtime.Callers to get the pc.
// This is solely for benchmarking the slowdown from runtime.Callers.
var IgnorePC = false
