package tagalign

import (
	"github.com/4meepo/tagalign"

	"github.com/flimzy/gopherlint/pkg/config"
	"github.com/flimzy/gopherlint/pkg/goanalysis"
)

func New(settings *config.TagAlignSettings) *goanalysis.Linter {
	var options []tagalign.Option

	if settings != nil {
		options = append(options, tagalign.WithAlign(settings.Align))

		if settings.Sort || len(settings.Order) > 0 {
			options = append(options, tagalign.WithSort(settings.Order...))
		}

		// Strict style will be applied only if Align and Sort are enabled together.
		if settings.Strict && settings.Align && settings.Sort {
			options = append(options, tagalign.WithStrictStyle())
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(tagalign.NewAnalyzer(options...)).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
