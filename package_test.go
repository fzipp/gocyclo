package gocyclo_test

import (
	"fmt"
	"github.com/gregoryv/gocyclo"
	"path/filepath"
	"strings"
	"testing"
)

var files []string

func init() {
	files = must(filepath.Glob("*.go"))
}

func TestAssert(t *testing.T) {
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
	result := gocyclo.Analyze(files)
	filtered := gocyclo.Filter(result, 1, -1)
	if len(filtered) != 1 {
		t.Fail()
	}
}

func TestAverage(t *testing.T) {
	result := gocyclo.Analyze(files)
	avg := gocyclo.Average(result)
	if avg > 3.0 {
		t.Errorf("%v", avg)
	}
}

func TestStat_String(t *testing.T) {
	result := gocyclo.Analyze([]string{"package_test.go"})
	out := result[0].String()
	if !strings.Contains(out, "package_test.go") {
		t.Error(out)
	}
}

// TestComplexity is an example of how to use it in your own code
func TestComplexity(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	max := 5
	result, ok := gocyclo.Assert(files, max)
	if !ok {
		for _, l := range result {
			fmt.Println(l)
		}
		t.Errorf("Exceeded maximum complexity %v", max)
	}
}

func must(result []string, err error) []string {
	if err != nil {
		panic(err)
	}
	return result
}
