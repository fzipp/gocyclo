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

func SortAndFilterStats(stats []Stat, top, over int) []Stat {
	sort.Sort(byComplexityDesc(stats))
	for i, stat := range stats {
		if i == top {
			return stats[:i]
		}
		if stat.Complexity <= over {
			return stats[:i]
		}
	}
	return stats
}

type byComplexityDesc []Stat

func (s byComplexityDesc) Len() int      { return len(s) }
func (s byComplexityDesc) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byComplexityDesc) Less(i, j int) bool {
	return s[i].Complexity >= s[j].Complexity
}
