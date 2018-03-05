package gocyclo_test

import (
	"github.com/gregoryv/gocyclo"
	"testing"
)

func TestAnalyze(t *testing.T) {
	gocyclo.Analyze([]string{})
}
