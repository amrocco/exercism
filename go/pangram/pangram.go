package pangram

import "regexp"
import "strings"

const testVersion = 1

//IsPangram determines if a string is a pangram or not
func IsPangram(sentence string) bool {
	letters := make(map[rune]bool)
	regex := regexp.MustCompile("[^a-zA-Z]+")

	sentence = strings.ToLower(sentence)
	sentence = regex.ReplaceAllString(sentence, "")

	for _, letter := range sentence {
		letters[letter] = true
	}

	if len(letters) == 26 {
		return true

	}

	return false
}
