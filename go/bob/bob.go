package bob

import "regexp"
import "strings"

const testVersion = 3

//Hey returns Bob's response to a given sentence
func Hey(sentence string) string {
	sentence = strings.TrimSpace(sentence)

	if len(sentence) == 0 {
		return "Fine. Be that way!"
	}

	regex := regexp.MustCompile("[^a-zA-Z]+")
	alphaSentence := regex.ReplaceAllString(sentence, "")

	if len(alphaSentence) != 0 && alphaSentence == strings.ToUpper(alphaSentence) {
		return "Whoa, chill out!"
	} else if sentence[len(sentence)-1:] == "?" {
		return "Sure."
	}

	return "Whatever."
}
