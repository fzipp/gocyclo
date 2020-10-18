package testdata

var noFuncLit = 1

var lit1 = func() {
}

var lit2 = func() {
	if true {
	}
}

var lit3 = func() {
	if true {
	}
	if true {
	}
}
