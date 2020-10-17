[![PkgGoDev](https://pkg.go.dev/badge/github.com/fzipp/gocyclo)](https://pkg.go.dev/github.com/fzipp/gocyclo)
[![Go Report Card](https://goreportcard.com/badge/github.com/fzipp/gocyclo)](https://goreportcard.com/report/github.com/fzipp/gocyclo)

Gocyclo calculates cyclomatic complexities of functions in Go source code.

The cyclomatic complexity of a function is calculated according to the
following rules:

     1 is the base complexity of a function
    +1 for each 'if', 'for', 'case', 'default', '&&' or '||'

To install, run

    $ go get github.com/fzipp/gocyclo/cmd/gocyclo

and put the resulting binary in one of your PATH directories if
`$GOPATH/bin` isn't already in your PATH.

Usage:

    $ gocyclo [<flag> ...] <Go file or directory> ...

Examples:

    $ gocyclo .
    $ gocyclo main.go
    $ gocyclo -top 10 src/
    $ gocyclo -over 25 docker
    $ gocyclo -avg .
    $ gocyclo -top 20 -ignore "_test|Godeps|vendor/" .

The output fields for each line are:

    <complexity> <package> <function> <file:row:column>

Individual functions can be ignored with a `gocyclo:ignore` directive:

    //gocyclo:ignore
    func f() {
        // ...
    }
    
    //gocyclo:ignore
    var g = func() {
    }
