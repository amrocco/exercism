package pangram

import "strings"

const testVersion = 1
const alpha = "abcdefghijklmnopqrstuvwxyz"

//IsPangram determines if a string is a pangram or not
func IsPangram(sentence string) bool {
	letters := make(map[rune]bool)
	sentence = strings.ToLower(sentence)

	for _, char := range sentence {
		if strings.Contains(alpha, string(char)) {
			letters[char] = true
		}
	}

	return len(letters) == 26
}
