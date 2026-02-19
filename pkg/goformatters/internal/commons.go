package internal

import "github.com/flimzy/gopherlint/pkg/logutils"

// FormatterLogger must be used only when the context logger is not available.
var FormatterLogger = logutils.NewStderrLog(logutils.DebugKeyFormatter)
