package testdata

func f1() {
}

func f2() {
	if true {
	}
}

func f2else() {
	if true {
	} else {
	}
}

func f3() {
	if true {
	}
	if true {
	}
}

func f3nested() {
	if true {
		if true {
		}
	}
}
