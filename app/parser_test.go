package app

import "testing"

func Test_getMaxPostCodeAndCount(t *testing.T) {
	postcodes := map[string]uint32{
		"10101": uint32(101),
		"10102": uint32(102),
		"10103": uint32(103),
	}

	expectedMax := uint32(103)
	expectedMaxPostcode := "10103"

	t.Run("testing getMaxPostCodeAndCount", func(t *testing.T) {
		max, maxPostcode := getMaxPostCodeAndCount(postcodes)
		if max != expectedMax {
			t.Errorf("got %d, want %d", max, expectedMax)
		}
		if maxPostcode != expectedMaxPostcode {
			t.Errorf("got %s, want %s", maxPostcode, expectedMaxPostcode)
		}
	})
}
