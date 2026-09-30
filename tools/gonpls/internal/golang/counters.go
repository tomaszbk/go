package golang

import "golang.org/x/telemetry/counter"

// Proposed counters for evaluating gopls extract, inline, and package rename. These counters
// increment when the user attempts to perform one of these operations,
// regardless of whether it succeeds.
var (
	countExtractFunction    = counter.New("gopls/extract:func")
	countExtractMethod      = counter.New("gopls/extract:method")
	countExtractVariable    = counter.New("gopls/extract:variable")
	countExtractVariableAll = counter.New("gopls/extract:variable-all")

	countInlineCall     = counter.New("gopls/inline:call")
	countInlineVariable = counter.New("gopls/inline:variable")

	countRenamePackage = counter.New("gopls/renamekind:package")
)
