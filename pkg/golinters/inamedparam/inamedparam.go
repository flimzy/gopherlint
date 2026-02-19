package inamedparam

import (
	"github.com/macabu/inamedparam"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.INamedParamSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			"skip-single-param": settings.SkipSingleParam,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(inamedparam.Analyzer).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
