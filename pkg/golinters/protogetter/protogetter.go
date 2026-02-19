package protogetter

import (
	"github.com/ghostiam/protogetter"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.ProtoGetterSettings) *goanalysis.Linter {
	var cfg protogetter.Config

	if settings != nil {
		cfg = protogetter.Config{
			SkipGeneratedBy:         settings.SkipGeneratedBy,
			SkipFiles:               settings.SkipFiles,
			SkipAnyGenerated:        settings.SkipAnyGenerated,
			ReplaceFirstArgInAppend: settings.ReplaceFirstArgInAppend,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(protogetter.NewAnalyzer(&cfg)).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
