Command gocyclo calculates cyclomatic complexities of functions in Go source code.

For more information on the metric refer to https://en.wikipedia.org/wiki/Cyclomatic_complexity.

To install, run

    $ go get github.com/fzipp/gocyclo

Usage:

    $ gocyclo [<flag> ...] <Go file or package> ...

Examples:

    $ gocyclo .
    $ gocyclo main.go
    $ gocyclo -top 10 src/
    $ gocyclo -over 25 docker
    $ gocyclo -avg .

The output fields for each line are:

    <complexity> <full function name> <file:row:column>
