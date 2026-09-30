package main

// Verify that we can import package "a" containing an inlineable
// function F that declares a local interface with a non-exported
// method f.
import _ "./a"

func main() {}
