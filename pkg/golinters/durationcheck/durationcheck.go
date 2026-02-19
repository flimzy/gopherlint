package durationcheck

import (
	"github.com/charithe/durationcheck"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(durationcheck.Analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
