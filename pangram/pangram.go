package pangram

import "strings"

const testVersion = 1

func IsPangram(i string) bool {
	alphabets := "abcdefghijklmnopqrstuvwxyz"
	alphabet_array := strings.Split(alphabets, "")

	for _, v := range alphabet_array {
		if strings.ContainsAny(strings.ToLower(i), v) == false {
			return false
		}
	}

	return true
}
