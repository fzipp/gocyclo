// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
)

func writeStats(w io.Writer, sortedStats []stat, top, over int) int {
	for i, stat := range sortedStats {
		if i == top {
			return i
		}
		if stat.Complexity <= over {
			return i
		}
		fmt.Fprintln(w, stat)
	}
	return len(sortedStats)
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
