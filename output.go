// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sort"
)

func sortAndFilterStats(stats []stat, top, over int) []stat {
	sort.Sort(byComplexityDesc(stats))
	i := 0
	for _, stat := range stats {
		if i == top {
			break
		}
		if stat.Complexity <= over {
			break
		}
		i++
	}
	return stats[:i]
}

func printStats(stats []stat) {
	for _, stat := range stats {
		fmt.Println(stat)
	}
}

func showAverage(stats []stat) {
	fmt.Printf("Average: %.3g\n", average(stats))
}

func average(stats []stat) float64 {
	return float64(sumTotal(stats)) / float64(len(stats))
}

func showTotal(stats []stat) {
	fmt.Printf("Total: %d\n", sumTotal(stats))
}

func sumTotal(stats []stat) uint64 {
	total := uint64(0)
	for _, s := range stats {
		total += uint64(s.Complexity)
	}
	return total
}
