package main

import (
	"cmd/internal/objabi"
	"cmd/internal/telemetry/counter"

	"golang.org/x/tools/go/analysis/suite/vet"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	// Keep consistent with cmd/fix/main.go!
	counter.Open()
	objabi.AddVersionFlag()
	counter.Inc("vet/invocations")

	unitchecker.Main(vet.Suite...) // (never returns)
}
