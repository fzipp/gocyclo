package testdata

type S struct{}

func (s S) m1() {
}

func (s S) m2() {
	if true {
	}
}

func (s *S) m1ptr() {
}

func (s *S) m2ptr() {
	if true {
	}
}
