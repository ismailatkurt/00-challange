package app

import (
	"bufio"
	"strconv"
)

// Whole parsing and matching/comparing logic lies here
func parseRecipesJson(scanner *bufio.Scanner, inputs Input) (CountPerPostcodeAndTime, map[string]uint16, map[string]uint32) {
	// Using uint16 since "The number of distinct recipe names is lower than 2K"
	recipeNames := make(map[string]uint16)

	// Using uint32 since "The number of distinct postcodes is lower than 1M"
	postcodes := make(map[string]uint32)

	countPerPostcodeAndTime := CountPerPostcodeAndTime{
		Postcode:      strconv.Itoa(int(inputs.Postcode)),
		From:          convertTo12HourFormat(inputs.From),
		To:            convertTo12HourFormat(inputs.To),
		DeliveryCount: 0, // will be incremented later in scan loop if postcode and from-to match
	}

	reader := Reader{s: scanner}
	for reader.s.Scan() { // this is reading the first "[" and ending of a recipe "}"
		// read opening brace line or first line after a recipe parsed/processed
		reader.s.Scan()
		t := reader.s.Text()
		// If end of file just break
		if t == "]" {
			break
		}

		postcode := reader.ReadPostcode()

		recipeName := reader.ReadRecipe()
		_, from, to := reader.ReadDelivery()

		if postcode == strconv.Itoa(int(inputs.Postcode)) {
			if inputs.From <= from && inputs.To >= to {
				countPerPostcodeAndTime.DeliveryCount++
			}
		}

		recipeNames[recipeName]++
		postcodes[postcode]++
	}

	return countPerPostcodeAndTime, recipeNames, postcodes
}

// Find Postcode with maximum delivery
func getMaxPostCodeAndCount(postcodes map[string]uint32) (uint32, string) {
	max := uint32(0)
	var maxPostCode string
	for p, c := range postcodes {
		if max < c {
			max = c
			maxPostCode = p
		}
	}

	return max, maxPostCode
}
