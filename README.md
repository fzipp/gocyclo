[![Build Status](https://travis-ci.org/gregoryv/gocyclo.svg?branch=master)](https://travis-ci.org/gregoryv/gocyclo)
[![codecov](https://codecov.io/gh/gregoryv/gocyclo/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/gocyclo)


Gocyclo calculates cyclomatic complexities of functions in Go source code.

The cyclomatic complexity of a function is calculated according to the
following rules:

     1 is the base complexity of a function
    +1 for each 'if', 'for', 'case', '&&' or '||'

To install, run

    $ go get -u github.com/gregoryv/gocyclo/...

Usage:

    $ gocyclo [<flag> ...] <Go file or directory> ...

Examples:

    $ gocyclo .
    $ gocyclo main.go
    $ gocyclo -top 10 src/
    $ gocyclo -over 5 *.go
    $ gocyclo -avg .

The output fields for each line are:

    <complexity> <package> <function> <file:row:column>


## As unit test

    func TestComplexity(t *testing.T) {
		files, err := filepath.Glob("*.go")
		if err != nil {
			t.Fatal(err)
		}
		max := 5
		result, ok := gocyclo.Assert(files, max)
		if !ok {
			for _, l := range result {
				fmt.Println(l)
			}
			t.Errorf("Exceeded maximum complexity %v", max)
		}
	}


