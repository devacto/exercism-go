package bob

import (
	"strings"
	"regexp"
)

const testVersion = 3

func Hey(input string) string {

	i := strings.TrimSpace(input)

	// Bob answers 'Whoa, chill out!' if you yell at him.

	// checks if input contains at least one character.
	b, _ := regexp.MatchString("[a-zA-Z]", i)
	if i == strings.ToUpper(i) && b == true {
		return "Whoa, chill out!"
	}

	// Bob answers 'Sure.' if you ask him a question.
	if strings.HasSuffix(i, "?") {
		return "Sure."
	}

	// Bob says 'Fine. Be that way!' if you address him without actually saying anything.
	if strings.TrimSpace(i) == "" {
		return "Fine. Be that way!"
	}

	// Bob answers 'Whatever.' to anything else.
	return "Whatever."
}