package goflow

import (
	"strconv"
)

func ParseInt(value string) int {
	if value == "" {
		return 0
	}
	val, _ := strconv.Atoi(value)
	return val
}

func IntString(value int) string {
	return strconv.Itoa(value)
}
