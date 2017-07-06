package leap

const testVersion = 3

// IsLeapYear returns true if the year is leap, otherwise returns false.
func IsLeapYear(y int) bool {
	return (y%4 == 0 && y%100 != 0) || y%400 == 0
}
