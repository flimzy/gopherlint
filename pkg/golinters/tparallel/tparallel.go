package tparallel

import (
	"github.com/moricho/tparallel"

	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(tparallel.Analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
