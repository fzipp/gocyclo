// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Gocyclo calculates the cyclomatic complexities of functions and
// methods in Go source code.
//
// Usage:
//      gocyclo [<flag> ...] <Go file or directory> ...
//
// Flags:
//      -over N   show functions with complexity > N only and
//                return exit code 1 if the output is non-empty
//      -top N    show the top N most complex functions only
//      -avg      show the average complexity
//
// The output fields for each line are:
// <complexity> <package> <function> <file:row:column>
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	. "github.com/fzipp/gocyclo/gocyclo"
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
<complexity> <package> <function> <file:row:column>
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
	log.SetFlags(0)
	log.SetPrefix("gocyclo: ")
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		usage()
	}

	stats := analyze(args)
	sort.Sort(ByComplexity(stats))
	written := writeStats(os.Stdout, stats)

	if *avg {
		ShowAverage(stats)
	}

	if *over > 0 && written > 0 {
		os.Exit(1)
	}
}

func writeStats(w io.Writer, sortedStats []Stat) int {
	for i, stat := range sortedStats {
		if i == *top {
			return i
		}
		if stat.Complexity <= *over {
			return i
		}
		fmt.Fprintln(w, stat)
	}
	return len(sortedStats)
}

func analyze(paths []string) []Stat {
	var stats []Stat
	for _, path := range paths {
		if isDir(path) {
			stats = AnalyzeDir(path, stats)
		} else {
			stats = AnalyzeFile(path, stats)
		}
	}
	return stats
}

func isDir(filename string) bool {
	fi, err := os.Stat(filename)
	return err == nil && fi.IsDir()
}
