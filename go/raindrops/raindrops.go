package raindrops

import (
	"bytes"
	"strconv"
)

const testVersion = 3

//Convert converts a number into a string based on the defined rules
func Convert(number int) string {
	var buffer bytes.Buffer
	result := ""

	if number%3 == 0 {
		buffer.WriteString("Pling")
	}
	if number%5 == 0 {
		buffer.WriteString("Plang")
	}
	if number%7 == 0 {
		buffer.WriteString("Plong")
	}

	result = buffer.String()

	if len(result) == 0 {
		result = strconv.Itoa(number)
	}

	return result
}
