// Package hamming
package hamming

import (
	"errors"
)

const testVersion = 6

func Distance(a, b string) (int, error) {
	var d = 0
	if len(a) != len(b) {
		return -1, errors.New("Strands need to be of the same length.")
	}
	for i := range a {
		if b[i] != a[i] {
			d++
		}
	}
	return d, nil
}
