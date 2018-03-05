package gocyclo_test

import (
	"fmt"
	"github.com/gregoryv/gocyclo"
	"path/filepath"
	"testing"
)

func TestAnalyze(t *testing.T) {
	files := must(filepath.Glob("*.go"))
	all := append(files, "cmd/")
	result, ok := gocyclo.Assert(all, 5)
	if !ok {
		for _, l := range result {
			fmt.Println(l)
		}
		t.Fail()
	}
	result, ok = gocyclo.Assert(all, -1)
	if ok {
		t.Error("Complexity should at least be 1")
	}
}

func TestFilter(t *testing.T) {
	files := must(filepath.Glob("*.go"))
	result := gocyclo.Analyze(files)
	filtered := gocyclo.Filter(result, 1, -1)
	if len(filtered) != 1 {
		t.Fail()
	}
}

func must(result []string, err error) []string {
	if err != nil {
		panic(err)
	}
	return result
}
