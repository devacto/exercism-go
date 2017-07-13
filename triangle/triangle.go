package triangle

import (
	"sort"
	"math"
)

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	// If any of the sides is not a number then not a triangle.
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}

	// If any of the sides is positive infinity then not a triangle.
	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return NaT
	}

	// If any of the sides is negative infinity then not a triangle.
	if math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) {
		return NaT
	}

	// If any of the sides is 0 then not a triangle.
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	// If fails triangle equality then not a triangle.
	s := []float64{a, b, c}
	sort.Float64s(s)
	if s[0] + s[1] < s[2] {
		return NaT
	}

	// If all three sides are the same then Equilateral.
	if a == b && b == c {
		return Equ
	}

	// If all three sides are not the same then Scalene
	if a != b && b != c && a != c {
		return Sca
	}

	// Anything else is Isosceles
	return Iso
}

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

// Pick values for the following identifiers used by the test program.
const (
	NaT Kind = iota // not a triangle - NaT to be 0
	Equ // equilateral - Equ to be 1
	Iso // isosceles - Iso to be 2
	Sca // scalene - Sca to be 3
)

// Organize your code for readability.
