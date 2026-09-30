package types2_test

import (
	"cmd/compile/internal/testimporter"
	"cmd/compile/internal/types2"
)

var imp = testimporter.NewImporter()

func defaultImporter() types2.Importer {
	return imp
}
