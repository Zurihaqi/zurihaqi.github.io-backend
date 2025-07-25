package utils

import (
	"errors"
	"strconv"
)

func StringToUint(s string) (uint, error) {
	val64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, errors.New("invalid uint string")
	}
	return uint(val64), nil
}
