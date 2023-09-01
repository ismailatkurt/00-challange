package main

import (
	"strconv"
)

func ConvertTo24HourFormat(s string) uint8 {
	f := s[len(s)-2:]
	n, _ := strconv.Atoi(s[:len(s)-2])

	if f == "AM" {
		return uint8(n)
	}

	return uint8(n + 12)
}

func ConvertTo12HourFormat(i uint8) string {
	if i < 12 {
		return strconv.Itoa(int(i)) + "AM"
	}

	return strconv.Itoa(int(i-12)) + "PM"
}
