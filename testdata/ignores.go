package testdata

//gocyclo:ignore
func ignoredFuncDecl() {}

//gocyclo:ignore
var ignoredFuncLit = func() {}

//gocyclo:skip
func notIgnoredUnknownDirective() {}

// gocyclo:ignore
func notIgnoredNotADirective() {}
