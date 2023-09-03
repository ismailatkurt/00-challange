package app

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func Test_getPostcodeFromLine(t *testing.T) {
	var tests = []struct {
		given    string
		expected string
	}{
		{`    "postcode": "10192",`, "10192"},
		{`"postcode": "10178",`, "10178"},
	}

	r := getReaderWithMockScanner("here is some mock string")

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.given, tt.expected)
		t.Run(testname, func(t *testing.T) {
			ans := r.getPostcodeFromLine(&tt.given)
			if ans != tt.expected {
				t.Errorf("got %s, want %s", ans, tt.expected)
			}
		})
	}
}

func Test_getRecipeFromLine(t *testing.T) {
	var tests = []struct {
		given    string
		expected string
	}{
		{`    "recipe": "Tex-Mex Tilapia",`, "Tex-Mex Tilapia"},
		{`"recipe": "Chicken Pineapple Quesadillas",`, "Chicken Pineapple Quesadillas"},
	}

	r := getReaderWithMockScanner("here is some mock string")

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.given, tt.expected)
		t.Run(testname, func(t *testing.T) {
			ans := r.getRecipeFromLine(&tt.given)
			if ans != tt.expected {
				t.Errorf("got %s, want %s", ans, tt.expected)
			}
		})
	}
}

func Test_getDeliveryFromLine(t *testing.T) {
	var tests = []struct {
		given    string
		expected string
	}{
		{`    "delivery": "Friday 11AM - 5PM"`, "Friday 11AM - 5PM"},
		{`"delivery": "Thursday 5AM - 7PM"`, "Thursday 5AM - 7PM"},
	}

	r := getReaderWithMockScanner("here is some mock string")

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.given, tt.expected)
		t.Run(testname, func(t *testing.T) {
			ans := r.getDeliveryFromLine(&tt.given)
			if ans != tt.expected {
				t.Errorf("got %s, want %s", ans, tt.expected)
			}
		})
	}
}

func TestReader_ReadPostcode(t *testing.T) {
	var tests = []struct {
		given    string
		expected string
	}{
		{`    "postcode": "10192",`, "10192"},
		{`"postcode": "10178",`, "10178"},
	}

	for _, tt := range tests {
		r := getReaderWithMockScanner(tt.given)

		testname := fmt.Sprintf("%s,%s", tt.given, tt.expected)
		t.Run(testname, func(t *testing.T) {
			ans := r.ReadPostcode()
			if ans != tt.expected {
				t.Errorf("got %s, want %s", ans, tt.expected)
			}
		})
	}
}

func TestReader_ReadRecipe(t *testing.T) {
	var tests = []struct {
		given    string
		expected string
	}{
		{`    "recipe": "Tex-Mex Tilapia",`, "Tex-Mex Tilapia"},
		{`"recipe": "Chicken Pineapple Quesadillas",`, "Chicken Pineapple Quesadillas"},
	}

	for _, tt := range tests {
		r := getReaderWithMockScanner(tt.given)

		testname := fmt.Sprintf("%s,%s", tt.given, tt.expected)
		t.Run(testname, func(t *testing.T) {
			ans := r.ReadRecipe()
			if ans != tt.expected {
				t.Errorf("got %s, want %s", ans, tt.expected)
			}
		})
	}
}

func TestReader_ReadDelivery(t *testing.T) {
	var tests = []struct {
		given    string
		expected struct {
			day  string
			from uint8
			to   uint8
		}
	}{
		{`    "delivery": "Friday 11AM - 5PM"`, struct {
			day  string
			from uint8
			to   uint8
		}{day: "Friday", from: 11, to: 17}},
		{`"delivery": "Thursday 5AM - 7PM"`, struct {
			day  string
			from uint8
			to   uint8
		}{day: "Thursday", from: 5, to: 19}},
	}

	for _, tt := range tests {
		r := getReaderWithMockScanner(tt.given)

		testname := fmt.Sprintf("%s", tt.given)
		t.Run(testname, func(t *testing.T) {
			day, from, to := r.ReadDelivery()
			if day != tt.expected.day {
				t.Errorf("got %s, want %s", day, tt.expected.day)
			}
			if from != tt.expected.from {
				t.Errorf("got %d, want %d", from, tt.expected.from)
			}
			if to != tt.expected.to {
				t.Errorf("got %d, want %d", to, tt.expected.to)
			}
		})
	}
}

func getReaderWithMockScanner(given string) Reader {
	mockReader := strings.NewReader(given)
	s := bufio.NewScanner(mockReader)
	r := Reader{s: s}

	return r
}
