package exptostd

import (
	"github.com/ldez/exptostd"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(exptostd.NewAnalyzer()).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
