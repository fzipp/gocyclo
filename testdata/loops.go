package testdata

func l2() {
	for {
	}
}

func l3() {
	for {
		for {
		}
	}
}
func l2range() {
	for range []struct{}{} {
	}
}

func l4() {
	for {
	}
	for true {
	}
	for ;; {
	}
}
