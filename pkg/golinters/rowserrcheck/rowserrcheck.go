package rowserrcheck

import (
	"github.com/jingyugao/rowserrcheck/passes/rowserr"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.RowsErrCheckSettings) *goanalysis.Linter {
	var pkgs []string

	if settings != nil {
		pkgs = settings.Packages
	}

	return goanalysis.
		NewLinterFromAnalyzer(rowserr.NewAnalyzer(pkgs...)).
		WithDesc("checks whether Rows.Err of rows is checked successfully").
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
