// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func analyze(paths []string) []stat {
	var stats []stat
	for _, path := range paths {
		if isDir(path) {
			stats = analyzeDir(path, stats)
		} else {
			stats = analyzeFile(path, stats)
		}
	}
	return stats
}

func isDir(filename string) bool {
	fi, err := os.Stat(filename)
	return err == nil && fi.IsDir()
}

func analyzeDir(dirname string, stats []stat) []stat {
	filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(path, ".go") {
			stats = analyzeFile(path, stats)
		}
		return err
	})
	return stats
}

func analyzeFile(fname string, stats []stat) []stat {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	analyzer := &fileAnalyzer{
		file:    f,
		fileSet: fset,
		stats:   stats,
	}
	analyzer.Analyze()
	return analyzer.stats
}

type fileAnalyzer struct {
	file    *ast.File
	fileSet *token.FileSet
	stats   []stat
}

func (a *fileAnalyzer) Analyze() {
	for _, declaration := range a.file.Decls {
		switch decl := declaration.(type) {
		case *ast.FuncDecl:
			a.addStatIfNotIgnored(decl, funcName(decl), decl.Doc)
		case *ast.GenDecl:
			for _, spec := range decl.Specs {
				valueSpec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				for _, value := range valueSpec.Values {
					funcLit, ok := value.(*ast.FuncLit)
					if !ok {
						continue
					}
					a.addStatIfNotIgnored(funcLit, valueSpec.Names[0].Name, decl.Doc)
				}
			}
		}
	}
}

func (a *fileAnalyzer) addStatIfNotIgnored(funcNode ast.Node, funcName string, doc *ast.CommentGroup) {
	if parseDirectives(doc).HasIgnore() {
		return
	}
	a.stats = append(a.stats, stat{
		PkgName:    a.file.Name.Name,
		FuncName:   funcName,
		Complexity: complexity(funcNode),
		Pos:        a.fileSet.Position(funcNode.Pos()),
	})
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
