package bob

import (
	"strings"
	"regexp"
)

const testVersion = 3

func Hey(input string) string {

	i := strings.TrimSpace(input)

	//He answers 'Whoa, chill out!' if you yell at him.
	b, _ := regexp.MatchString("[a-zA-Z]", i)

	// if input are all upper case and input contains at least one character then shouting.
	if i == strings.ToUpper(i) && b == true {
		return "Whoa, chill out!"
	}

	//Bob answers 'Sure.' if you ask him a question.
	if strings.HasSuffix(i, "?") {
		return "Sure."
	}

	//He says 'Fine. Be that way!' if you address him without actually saying anything.
	if strings.TrimSpace(i) == "" {
		return "Fine. Be that way!"
	}

	//He answers 'Whatever.' to anything else.
	return "Whatever."
}