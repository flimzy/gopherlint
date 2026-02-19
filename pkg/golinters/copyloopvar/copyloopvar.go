package copyloopvar

import (
	"github.com/karamaru-alpha/copyloopvar"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.CopyLoopVarSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			"check-alias": settings.CheckAlias,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(copyloopvar.NewAnalyzer()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
