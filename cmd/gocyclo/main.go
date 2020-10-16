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
//      -total    show the total complexity
//
// The output fields for each line are:
// <complexity> <package> <function> <file:row:column>
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fzipp/gocyclo"
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
        -total    show the total complexity for all functions

The output fields for each line are:
<complexity> <package> <function> <file:row:column>
`

func main() {
	over := flag.Int("over", 0, "show functions with complexity > N only")
	top := flag.Int("top", -1, "show the top N most complex functions only")
	avg := flag.Bool("avg", false, "show the average complexity")
	total := flag.Bool("total", false, "show the total complexity")

	log.SetFlags(0)
	log.SetPrefix("gocyclo: ")
	flag.Usage = usage
	flag.Parse()
	paths := flag.Args()
	if len(paths) == 0 {
		usage()
	}

	allStats := gocyclo.Analyze(paths)
	shownStats := gocyclo.SortAndFilterStats(allStats, *top, *over)
	printStats(shownStats)

	if *avg {
		printAverage(allStats)
	}

	if *total {
		printTotal(allStats)
	}

	if *over > 0 && len(shownStats) > 0 {
		os.Exit(1)
	}
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, usageDoc)
	os.Exit(2)
}
