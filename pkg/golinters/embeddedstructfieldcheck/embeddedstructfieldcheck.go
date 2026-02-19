package embeddedstructfieldcheck

import (
	"github.com/manuelarte/embeddedstructfieldcheck/analyzer"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.EmbeddedStructFieldCheckSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			analyzer.ForbidMutexCheck: settings.ForbidMutex,
			analyzer.EmptyLineCheck:   settings.EmptyLine,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer.NewAnalyzer()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
