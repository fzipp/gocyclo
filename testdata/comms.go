package testdata

var ch chan struct{}

func comm2() {
	select {
	case <-ch:
	}
}

func comm2default() {
	select {
	case <-ch:
	default:
	}
}

func comm3() {
	select {
	case <-ch:
	case <-ch:
	}
}

func comm3default() {
	select {
	case <-ch:
	case <-ch:
	default:
	}
}

func comm3nested() {
	select {
	case <-ch:
		select {
		case <-ch:
		}
	}
}
