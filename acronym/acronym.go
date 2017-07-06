package acronym

import (
	"strings"
)

const testVersion = 3

func Abbreviate(i string) string {
	a := strings.FieldsFunc(i, vsplit)
	r := ""
	for _, v := range a {
		r += strings.ToUpper(string(v[0]))
	}
	return r
}

func vsplit(r rune) bool {
	return r == ' ' || r == '-'
}
