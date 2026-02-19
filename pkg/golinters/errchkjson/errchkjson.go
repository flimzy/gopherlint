package errchkjson

import (
	"github.com/breml/errchkjson"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.ErrChkJSONSettings) *goanalysis.Linter {
	cfg := map[string]any{
		"omit-safe": true,
	}

	if settings != nil {
		cfg = map[string]any{
			"omit-safe":          !settings.CheckErrorFreeEncoding,
			"report-no-exported": settings.ReportNoExported,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(errchkjson.NewAnalyzer()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
