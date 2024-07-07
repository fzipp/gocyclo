// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Gocyclo calculates the cyclomatic complexities of functions and
// methods in Go source code.
//
// Usage:
//
//	gocyclo [<flag> ...] <Go file or directory> ...
//
// Flags:
//
//	-over N                         show functions with complexity > N only and
//	                                return exit code 1 if the output is non-empty
//	-top N                          show the top N most complex functions only
//	-avg, -avg-short                show the average complexity;
//	                                the short option prints the value without a label
//	-centile, -centile-short K      show K-th percentile (1 <= K <= 99)
//	                                the short option prints the value without a label
//	-ignore REGEX                   exclude files matching the given regular expression
//
// The output fields for each line are:
// <complexity> <package> <function> <file:line:column>
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/fzipp/gocyclo"
)

const usageDoc = `Calculate cyclomatic complexities of Go functions.
Usage:
    gocyclo [flags] <Go file or directory> ...

Flags:
    -over N                         show functions with complexity > N only and
                                    return exit code 1 if the output is non-empty
    -top N                          show the top N most complex functions only
    -avg, -avg-short                show the average complexity;
                                    the short option prints the value without a label
    -centile, -centile-short K      show K-th percentile (1 <= K <= 99)
                                    the short option prints the value without a label
    -ignore REGEX                   exclude files matching the given regular expression

The output fields for each line are:
<complexity> <package> <function> <file:line:column>
`

func main() {
	over := flag.Int("over", 0, "show functions with complexity > N only")
	top := flag.Int("top", -1, "show the top N most complex functions only")
	avg := flag.Bool("avg", false, "show the average complexity")
	avgShort := flag.Bool("avg-short", false, "show the average complexity without a label")
	centile := flag.Int("centile", 0, "show K-th percentile (1 <= K <= 99)")
	centileShort := flag.Int("centile-short", 0, "show K-th percentile (1 <= K <= 99)")
	ignore := flag.String("ignore", "", "exclude files matching the given regular expression")

	log.SetFlags(0)
	log.SetPrefix("gocyclo: ")
	flag.Usage = usage
	flag.Parse()
	paths := flag.Args()
	if len(paths) == 0 {
		usage()
	}

	allStats := gocyclo.Analyze(paths, regex(*ignore))
	shownStats := allStats.SortAndFilter(*top, *over)

	printStats(shownStats)
	if *avg || *avgShort {
		printAverage(allStats, *avgShort)
	}

	centileLabel := true
	centileValue := 0
	if *centile > 0 && *centileShort > 0 {
		// centile and centile-short are mutually exclusive
		usage()
	} else if *centile > 0 {
		centileLabel = true
		centileValue = *centile
	} else if *centileShort > 0 {
		centileLabel = false
		centileValue = *centileShort
	}
	if centileValue > 99 {
		usage()
	} else if centileValue > 0 {
		printCentile(allStats.SortAndFilter(-1, 0), centileValue, centileLabel)
	}

	if *over > 0 && len(shownStats) > 0 {
		os.Exit(1)
	}
}

func regex(expr string) *regexp.Regexp {
	if expr == "" {
		return nil
	}
	re, err := regexp.Compile(expr)
	if err != nil {
		log.Fatal(err)
	}
	return re
}

func printStats(s gocyclo.Stats) {
	for _, stat := range s {
		fmt.Println(stat)
	}
}

func printAverage(s gocyclo.Stats, short bool) {
	if !short {
		fmt.Print("Average: ")
	}
	fmt.Printf("%.3g\n", s.AverageComplexity())
}

func printCentile(s gocyclo.Stats, centile int, label bool) {
	res, err := s.Percentile(centile)
	if err != nil {
		log.Fatal(err)
	}

	indicators := []string{"th", "st", "nd", "rd", "th", "th", "th", "th", "th", "th"}
	if label {
		indicator := "th"
		if centile < 11 || centile > 13 {
			// 11th not 11st, same to 12 and 13
			indicator = indicators[centile%10]
		}
		fmt.Printf("%d%s-percentile: ", centile, indicator)
	}

	fmt.Printf("%d\n", res)
}

func usage() {
	_, _ = fmt.Fprint(os.Stderr, usageDoc)
	os.Exit(2)
}
