// conpack/main_test.go

package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestUsage(t *testing.T) {
	// Capture the output of the usage function
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	usage()

	w.Close()
	os.Stdout = old

	var buf strings.Builder
	_, _ = io.Copy(&buf, r)

	output := buf.String()

	expected := "Usage: \n  conpack [-p|--package <package_name>] [-r|--runtime <runtime>]"
	if !strings.Contains(output, expected) {
		t.Errorf("usage() = %q, want %q", output, expected)
	}
}
