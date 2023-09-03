package app

import (
	"bufio"
	"strconv"
)

func parseRecipesJson(scanner *bufio.Scanner, inputs Input) (CountPerPostcodeAndTime, map[string]uint16, map[string]uint32) {
	recipeNames := make(map[string]uint16)
	postcodes := make(map[string]uint32)
	countPerPostcodeAndTime := CountPerPostcodeAndTime{
		Postcode:      strconv.Itoa(int(inputs.Postcode)),
		From:          convertTo12HourFormat(inputs.From),
		To:            convertTo12HourFormat(inputs.To),
		DeliveryCount: 0,
	}

	reader := Reader{s: scanner}
	for reader.s.Scan() {
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
