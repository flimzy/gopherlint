package whitespace

import (
	"github.com/ultraware/whitespace"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.WhitespaceSettings) *goanalysis.Linter {
	var wsSettings whitespace.Settings
	if settings != nil {
		wsSettings = whitespace.Settings{
			MultiIf:   settings.MultiIf,
			MultiFunc: settings.MultiFunc,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(whitespace.NewAnalyzer(&wsSettings)).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
