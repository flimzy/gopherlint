package commands

import (
	"bytes"
	"testing"
)

func TestPrintVersionShouldReportItselfAsGopherlint(t *testing.T) {
	t.Parallel()

	info := BuildInfo{Version: "1.27.2", GoVersion: "go1.27.3", Commit: "abc123", Date: "2026-11-17"}
	want := "gopherlint has version 1.27.2 built with go1.27.3 from abc123 on 2026-11-17\n"

	var buf bytes.Buffer
	if err := printVersion(&buf, info); err != nil {
		t.Fatal(err)
	}

	if got := buf.String(); got != want {
		t.Errorf("unexpected output:\n got: %q\nwant: %q", got, want)
	}
}
