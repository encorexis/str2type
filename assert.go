package str2type

import "strconv"

func AssertBool(str ...string) bool {
	for _, s := range str {
		_, err := strconv.ParseBool(s)
		if err != nil {
			return false
		}
	}
	return true
}

func AssertInt(str ...string) bool {
	for _, s := range str {
		_, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
	}
	return true
}

func AssertFloat(str ...string) bool {
	for _, s := range str {
		_, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return false
		}
	}
	return true
}
