package nolintlint

import (
	"testing"

	"github.com/flimzy/gopherlint/test/testshared/integration"
)

func TestFromTestdata(t *testing.T) {
	integration.RunTestdata(t)
}

func TestFix(t *testing.T) {
	integration.RunFix(t)
}

func TestFixPathPrefix(t *testing.T) {
	integration.RunFixPathPrefix(t)
}
