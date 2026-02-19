package nakedret

import (
	"github.com/alexkohler/nakedret/v2"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.NakedretSettings) *goanalysis.Linter {
	cfg := &nakedret.NakedReturnRunner{}

	if settings != nil {
		// SkipTestFiles is intentionally ignored => should be managed with `linters.exclusions.rules`.
		cfg.MaxLength = settings.MaxFuncLines
	}

	return goanalysis.
		NewLinterFromAnalyzer(nakedret.NakedReturnAnalyzer(cfg)).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
