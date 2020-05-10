package maths

import "strconv"

func IsNumber(num string) bool {
	_, err := strconv.Atoi(num)
	return err == nil
}
