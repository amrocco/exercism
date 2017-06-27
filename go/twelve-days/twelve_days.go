package twelve

import (
	"fmt"
	"strings"
)

const testVersion = 1

//Song returns all twelve verses of the Twelve Days of Christmas
func Song() string {
	var song []string

	for i := 1; i <= 12; i++ {
		song = append(song, Verse(i))
	}

	song = append(song, "")

	return strings.Join(song, "\n")
}

//Verse returns the corresponding verse from the Twelve Days of Christmas
func Verse(number int) string {
	var gifts []string
	verse := ""

	for i := number; i > 0; i-- {
		gifts = append(gifts, gift(i))
	}

	verse = strings.Join(gifts, ", ")

	if number == 1 {
		verse = verse[4:]
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me, %s.", ordinalNumber(number), verse)
}

func gift(number int) string {
	switch number {
	case 1:
		return "and a Partridge in a Pear Tree"
	case 2:
		return "two Turtle Doves"
	case 3:
		return "three French Hens"
	case 4:
		return "four Calling Birds"
	case 5:
		return "five Gold Rings"
	case 6:
		return "six Geese-a-Laying"
	case 7:
		return "seven Swans-a-Swimming"
	case 8:
		return "eight Maids-a-Milking"
	case 9:
		return "nine Ladies Dancing"
	case 10:
		return "ten Lords-a-Leaping"
	case 11:
		return "eleven Pipers Piping"
	case 12:
		return "twelve Drummers Drumming"
	}

	return "Input must be in the range of 1..12"
}

func ordinalNumber(number int) string {
	switch number {
	case 1:
		return "first"
	case 2:
		return "second"
	case 3:
		return "third"
	case 4:
		return "fourth"
	case 5:
		return "fifth"
	case 6:
		return "sixth"
	case 7:
		return "seventh"
	case 8:
		return "eighth"
	case 9:
		return "ninth"
	case 10:
		return "tenth"
	case 11:
		return "eleventh"
	case 12:
		return "twelfth"
	}

	return "Input must be in the range of 1..12"
}
