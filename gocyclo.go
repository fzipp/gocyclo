// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command gocyclo calculates the cyclomatic complexities of functions
// in Go source code.
//
// Usage:
//      gocyclo [<flag> ...] <Go file or packages> ...
//
// Flags:
//      -over N   show functions with complexity > N only and
//                return exit code 1 if the output is non-empty
//      -top N    show the top N most complex functions only
//      -avg      show the average complexity
//
// The output fields for each line are:
// <complexity> <full function name> <file:row:column>
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

const usageDoc = `Calculate cyclomatic complexities of Go functions.
Usage:
        gocyclo [flags] <Go file or directory> ...

Flags:
        -over N   show functions with complexity > N only and
                  return exit code 1 if the set is non-empty
        -top N    show the top N most complex functions only
        -avg      show the average complexity over all functions,
                  not depending on whether -over or -top are set

The output fields for each line are:
<complexity> <full function name> <file:row:column>
`

func usage() {
	fmt.Fprintf(os.Stderr, usageDoc)
	os.Exit(2)
}

var (
	over = flag.Int("over", 0, "show functions with complexity > N only")
	top  = flag.Int("top", -1, "show the top N most complex functions only")
	avg  = flag.Bool("avg", false, "show the average complexity")
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	conf := loader.Config{}
	args, err := conf.FromArgs(args, false)
	if err != nil {
		return err
	}

	lprog, err := conf.Load()
	if err != nil {
		return err
	}

	prog := ssautil.CreateProgram(lprog, ssa.NaiveForm)
	stats := analyze(lprog, prog)
	sort.Sort(byComplexity(stats))
	written := writeStats(os.Stdout, stats)

	if *avg {
		fmt.Printf("Average: %.3g\n", average(stats))
	}

	if *over > 0 && written > 0 {
		os.Exit(1)
	}
	return nil
}

type stat struct {
	Func       *ssa.Function
	Complexity int
}

func (s stat) String() string {
	pos := s.Func.Prog.Fset.Position(s.Func.Pos())
	return fmt.Sprintf("%d %s %s", s.Complexity, s.Func.String(), pos)
}

func analyze(lprog *loader.Program, prog *ssa.Program) []stat {
	var stats []stat
	for _, lpkg := range lprog.InitialPackages() {
		pkg := prog.Package(lpkg.Pkg)
		pkg.Build()
		for _, m := range pkg.Members {
			if fn, ok := m.(*ssa.Function); ok {
				stats = append(stats, stat{
					Func:       fn,
					Complexity: complexity(fn),
				})
			}
		}
	}
	return stats
}

// complexity calculates the cyclomatic complexity of a function.
func complexity(fn *ssa.Function) int {
	// https://en.wikipedia.org/wiki/Cyclomatic_complexity
	// The complexity M for a function is defined as
	// M = E − N + 2
	// where
	//
	// E = the number of edges of the graph.
	// N = the number of nodes of the graph.
	edges := 0
	for _, b := range fn.Blocks {
		edges += len(b.Succs)
	}
	return edges - len(fn.Blocks) + 2
}

func writeStats(w io.Writer, sortedStats []stat) int {
	written := 0
	for i, stat := range sortedStats {
		if i == *top {
			break
		}
		if stat.Complexity <= *over {
			break
		}
		written++
		fmt.Fprintln(w, stat)
	}
	return written
}

func average(stats []stat) float64 {
	total := 0
	for _, s := range stats {
		total += s.Complexity
	}
	return float64(total) / float64(len(stats))
}

type byComplexity []stat

func (s byComplexity) Len() int      { return len(s) }
func (s byComplexity) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byComplexity) Less(i, j int) bool {
	return s[i].Complexity >= s[j].Complexity
}
