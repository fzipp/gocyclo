// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/fzipp/gocyclo"
)

func printStats(stats []gocyclo.Stat) {
	for _, stat := range stats {
		fmt.Println(stat)
	}
}

func showAverage(stats []gocyclo.Stat) {
	fmt.Printf("Average: %.3g\n", average(stats))
}

func showTotal(stats []gocyclo.Stat) {
	fmt.Printf("Total: %d\n", sumTotal(stats))
}

func average(stats []gocyclo.Stat) float64 {
	return float64(sumTotal(stats)) / float64(len(stats))
}

func sumTotal(stats []gocyclo.Stat) uint64 {
	total := uint64(0)
	for _, s := range stats {
		total += uint64(s.Complexity)
	}
	return total
}
