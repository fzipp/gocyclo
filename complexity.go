// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gocyclo

import (
	"go/ast"
	"go/token"
)

// Complexity calculates the cyclomatic complexity of a function.
func Complexity(fn ast.Node) int {
	v := complexityVisitor{}
	ast.Walk(&v, fn)
	return v.complexity
}

type complexityVisitor struct {
	// complexity is the cyclomatic complexity
	complexity int
}

// Visit implements the ast.Visitor interface.
func (v *complexityVisitor) Visit(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.FuncDecl, *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.CaseClause, *ast.CommClause:
		v.complexity++
	case *ast.BinaryExpr:
		if n.Op == token.LAND || n.Op == token.LOR {
			v.complexity++
		}
	}
	return v
}
