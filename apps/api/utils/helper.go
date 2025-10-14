package utils

import "strconv"

func ToUint(s string) (uint, error) {
	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}