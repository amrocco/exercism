package gigasecond

import "time"

const testVersion = 4
const gigasecond = 1000000000

//AddGigasecond adds a gigasecond to a given time
func AddGigasecond(addend time.Time) time.Time {
	gigasecondDuration := (gigasecond * 1000) * time.Millisecond
	sum := addend.Add(gigasecondDuration)
	return sum
}
