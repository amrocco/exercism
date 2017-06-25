package leap

const testVersion = 3

//IsLeapYear is a function that determines if a calendar year is a leap year or not
func IsLeapYear(year int) bool {
	return year%4 == 0 && !(year%100 == 0 && year%400 != 0)
}
