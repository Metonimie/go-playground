package leap

const testVersion = 3

func IsLeapYear(year int) bool {

	if year % 4 == 0 { // True for every year that is divisible by 4
		if year % 400 == 0 { // If the year is divisible by 400 it surely is a leap year.
			return true
		}
		if year % 100 == 0 { // Unless it's not divisible by 400 and is divisible by 100
			return false
		}
		return true
	}

	return false
}
