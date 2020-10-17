// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gocyclo

import (
	"fmt"
	"go/token"
	"sort"
)

type Stat struct {
	PkgName    string
	FuncName   string
	Complexity int
	Pos        token.Position
}

func (s Stat) String() string {
	return fmt.Sprintf("%d %s %s %s", s.Complexity, s.PkgName, s.FuncName, s.Pos)
}

type Stats []Stat

func (s Stats) AverageComplexity() float64 {
	return float64(s.TotalComplexity()) / float64(len(s))
}

func (s Stats) TotalComplexity() uint64 {
	total := uint64(0)
	for _, stat := range s {
		total += uint64(stat.Complexity)
	}
	return total
}

func (s Stats) SortAndFilter(top, over int) Stats {
	result := make(Stats, len(s))
	copy(result, s)
	sort.Sort(byComplexityDesc(result))
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

type byComplexityDesc Stats

func (s byComplexityDesc) Len() int      { return len(s) }
func (s byComplexityDesc) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byComplexityDesc) Less(i, j int) bool {
	return s[i].Complexity >= s[j].Complexity
}
