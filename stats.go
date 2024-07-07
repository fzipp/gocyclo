// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gocyclo

import (
	"fmt"
	"go/token"
	"sort"
)

// Stat holds the cyclomatic complexity of a function, along with its package
// and and function name and its position in the source code.
type Stat struct {
	PkgName    string
	FuncName   string
	Complexity int
	Pos        token.Position
}

// String formats the cyclomatic complexity information of a function in
// the following format: "<complexity> <package> <function> <file:line:column>"
func (s Stat) String() string {
	return fmt.Sprintf("%d %s %s %s", s.Complexity, s.PkgName, s.FuncName, s.Pos)
}

// Stats hold the cyclomatic complexities of many functions.
type Stats []Stat

// AverageComplexity calculates the average cyclomatic complexity of the
// cyclomatic complexities in s.
func (s Stats) AverageComplexity() float64 {
	return float64(s.TotalComplexity()) / float64(len(s))
}

// TotalComplexity calculates the total sum of all cyclomatic
// complexities in s.
func (s Stats) TotalComplexity() uint64 {
	total := uint64(0)
	for _, stat := range s {
		total += uint64(stat.Complexity)
	}
	return total
}

// SortAndFilter sorts the cyclomatic complexities in s in descending order
// and returns a slice of s limited to the 'top' N entries with a cyclomatic
// complexity greater than 'over'. If 'top' is negative, i.e. -1, it does
// not limit the result. If 'over' is <= 0 it does not limit the result either,
// because a function has a base cyclomatic complexity of at least 1.
func (s Stats) SortAndFilter(top, over int) Stats {
	result := make(Stats, len(s))
	copy(result, s)
	sort.Stable(byComplexityDesc(result))
	for i, stat := range result {
		if i == top {
			return result[:i]
		}
		if stat.Complexity <= over {
			return result[:i]
		}
	}
	return result
}

// Percentile return k-th percentile, s needs to be sorted in decending
// order and k needs to be [1,99], otherwise error will be returned
func (s Stats) Percentile(k int) (int, error) {
	if !sort.SliceIsSorted(s, func(i, j int) bool { return s[i].Complexity > s[j].Complexity }) {
		return -1, fmt.Errorf("stats needs to be sorted before calculating k-th percentile")
	}
	if k < 1 || k > 99 {
		return -1, fmt.Errorf("k needs to be [1,99]")
	}
	loc := (100-k)*len(s)/100 - 1
	if loc < 0 {
		loc = 0
	}
	return s[loc].Complexity, nil
}

type byComplexityDesc Stats

func (s byComplexityDesc) Len() int      { return len(s) }
func (s byComplexityDesc) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byComplexityDesc) Less(i, j int) bool {
	return s[i].Complexity >= s[j].Complexity
}
