package clock

import "fmt"

const testVersion = 4
const minutesInDay = 1440

//Clock is a representation in minutes of a 24 hour clock
type Clock struct {
	minutes int
}

//New is the constructor for the Clock struct
func New(hour, minutes int) Clock {
	minutes = minutes + (hour * 60)
	clock := Clock{0}
	clock = clock.Add(minutes)
	return clock
}

//String returns the clock's time as a string in %H:%M format
func (clock Clock) String() string {
	hour := (clock.minutes / 60) % 24
	minutes := clock.minutes % 60
	return fmt.Sprintf("%02d:%02d", hour, minutes)
}

//Add adds n minutes to the clock's time
func (clock Clock) Add(minutes int) Clock {
	clock.minutes = (clock.minutes + minutes) % minutesInDay

	if clock.minutes < 0 {
		clock.minutes = minutesInDay + clock.minutes
	}

	return clock
}
