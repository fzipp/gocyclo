// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gocyclo

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func AnalyzeFile(fname string, stats []Stat) ([]Stat, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fname, nil, 0)
	if err != nil {
		return nil, err
	}
	return buildStats(f, fset, stats), nil
}

func AnalyzeDir(dirname string, stats []Stat) []Stat {
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(path, ".go") {
			stats, err = AnalyzeFile(path, stats)
			if err != nil {
				return err
			}
		}
		return err
	})
	return stats, err
}

func ShowAverage(stats []Stat) {
	fmt.Printf("Average: %.3g\n", Average(stats))
}

func Average(stats []Stat) float64 {
	total := 0
	for _, s := range stats {
		total += s.Complexity
	}
	return float64(total) / float64(len(stats))
}

type Stat struct {
	PkgName    string
	FuncName   string
	Complexity int
	Pos        token.Position
}

func (s Stat) String() string {
	return fmt.Sprintf("%d %s %s %s", s.Complexity, s.PkgName, s.FuncName, s.Pos)
}

type ByComplexity []Stat

func (s ByComplexity) Len() int      { return len(s) }
func (s ByComplexity) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByComplexity) Less(i, j int) bool {
	return s[i].Complexity >= s[j].Complexity
}

func buildStats(f *ast.File, fset *token.FileSet, stats []Stat) []Stat {
	for _, decl := range f.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			stats = append(stats, Stat{
				PkgName:    f.Name.Name,
				FuncName:   funcName(fn),
				Complexity: complexity(fn),
				Pos:        fset.Position(fn.Pos()),
			})
		}
	}
	return stats
}

// funcName returns the name representation of a function or method:
// "(Type).Name" for methods or simply "Name" for functions.
func funcName(fn *ast.FuncDecl) string {
	if fn.Recv != nil {
		if fn.Recv.NumFields() > 0 {
			typ := fn.Recv.List[0].Type
			return fmt.Sprintf("(%s).%s", recvString(typ), fn.Name)
		}
	}
	return fn.Name.Name
}

// recvString returns a string representation of recv of the
// form "T", "*T", or "BADRECV" (if not a proper receiver type).
func recvString(recv ast.Expr) string {
	switch t := recv.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + recvString(t.X)
	}
	return "BADRECV"
}

// complexity calculates the cyclomatic complexity of a function.
func complexity(fn *ast.FuncDecl) int {
	v := complexityVisitor{}
	ast.Walk(&v, fn)
	return v.Complexity
}

type complexityVisitor struct {
	// Complexity is the cyclomatic complexity
	Complexity int
}

// Visit implements the ast.Visitor interface.
func (v *complexityVisitor) Visit(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.FuncDecl, *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.CaseClause, *ast.CommClause:
		v.Complexity++
	case *ast.BinaryExpr:
		if n.Op == token.LAND || n.Op == token.LOR {
			v.Complexity++
		}
	}
	return v
}
