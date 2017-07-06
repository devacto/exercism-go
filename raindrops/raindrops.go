package raindrops

import "strconv"

const testVersion = 3

func Convert(n int) string {
	s := ""
	if n % 3 == 0 {
		s += "Pling"
	}
	if n % 5 == 0 {
		s += "Plang"
	}
	if n % 7 == 0 {
		s += "Plong"
	}
	if s == "" {
		s = strconv.Itoa(n)
	}
	return s
}

// Don't forget the test program has a benchmark too.
// How fast does your Convert convert?
