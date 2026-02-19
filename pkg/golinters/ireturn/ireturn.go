package ireturn

import (
	"strings"

	"github.com/butuzov/ireturn/analyzer"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.IreturnSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			"allow":    strings.Join(settings.Allow, ","),
			"reject":   strings.Join(settings.Reject, ","),
			"nonolint": true,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer.NewAnalyzer()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
