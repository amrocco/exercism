package acronym

import (
	"bytes"
	"regexp"
	"strings"
)

const testVersion = 3

//Abbreviate takes a phrase and returns its acronym
func Abbreviate(phrase string) string {
	var buffer bytes.Buffer
	words := regexp.MustCompile("[^a-zA-Z0-9_]+").Split(phrase, -1)

	for _, word := range words {
		firstLetter := string(word[0])
		firstLetter = strings.ToUpper(firstLetter)
		buffer.WriteString(firstLetter)
	}

	return buffer.String()
}
