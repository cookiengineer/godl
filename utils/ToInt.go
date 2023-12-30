package utils

import "strconv"

func ToInt(value string) int {

	number, err := strconv.ParseInt(value, 10, 64)

	if err == nil {
		return int(number)
	}

	return 0

}
