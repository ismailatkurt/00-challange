package app

import (
	"fmt"
	"testing"
)

func Test_convertTo24HourFormat(t *testing.T) {
	var tests = []struct {
		given    string
		expected uint8
	}{
		{"4AM", 4},
		{"12AM", 0},
		{"1PM", 13},
		{"12PM", 12},
		{"9PM", 21},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%d", tt.given, tt.expected)
		t.Run(testname, func(t *testing.T) {
			ans := convertTo24HourFormat(tt.given)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}

func Test_convertTo12HourFormat(t *testing.T) {
	var tests = []struct {
		given    uint8
		expected string
	}{
		{4, "4AM"},
		{0, "12AM"},
		{13, "1PM"},
		{12, "12PM"},
		{21, "9PM"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%s", tt.given, tt.expected)
		t.Run(testname, func(t *testing.T) {
			ans := convertTo12HourFormat(tt.given)
			if ans != tt.expected {
				t.Errorf("got %s, want %s", ans, tt.expected)
			}
		})
	}
}
