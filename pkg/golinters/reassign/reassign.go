package reassign

import (
	"fmt"
	"strings"

	"github.com/curioswitch/go-reassign"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.ReassignSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil && len(settings.Patterns) > 0 {
		cfg = map[string]any{
			reassign.FlagPattern: fmt.Sprintf("^(%s)$", strings.Join(settings.Patterns, "|")),
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(reassign.NewAnalyzer()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
