package testdata

func c1() {
}

func c2() {
	switch 1 {
	case 1:
	}
}

func c2default() {
	switch 1 {
	case 1:
	default:
	}
}

func c2multi() {
	switch 1 {
	case 1, 2:
	}
}

func c3() {
	switch 1 {
	case 1:
	case 2:
	}
}

func c3default() {
	switch 1 {
	case 1:
	case 2:
	default:
	}
}

func c3nested() {
	switch 1 {
	case 1:
		switch 2 {
		case 1:
		}
	}
}
