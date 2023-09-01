package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type CountPerRecipe struct {
	Recipe string `json:"recipe"`
	Count  uint   `json:"count"`
}

type BusiestPostcode struct {
	Postcode      string `json:"postcode"`
	DeliveryCount uint   `json:"delivery_count"`
}

type CountPerPostcodeAndTime struct {
	Postcode      string `json:"postcode"`
	From          string `json:"from"`
	To            string `json:"to"`
	DeliveryCount uint   `json:"delivery_count"`
}

type Result struct {
	UniqueRecipeCount       uint                    `json:"unique_recipe_count"`
	CountPerRecipe          []CountPerRecipe        `json:"count_per_recipe"`
	BusiestPostcode         BusiestPostcode         `json:"busiest_postcode"`
	CountPerPostcodeAndTime CountPerPostcodeAndTime `json:"count_per_postcode_and_time"`
	MatchByName             []string                `json:"match_by_name"`
}

func Process() *Result {
	inputs := getInputs()
	f := openFile()
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			_ = fmt.Errorf("could not close the file: %s", err.Error())
		}
	}(f)

	recipeNames := make(map[string]int)
	postcodes := make(map[string]int)
	countPerPostcodeAndTime := CountPerPostcodeAndTime{
		Postcode:      strconv.Itoa(int(inputs.Postcode)),
		From:          ConvertTo12HourFormat(inputs.From),
		To:            ConvertTo12HourFormat(inputs.To),
		DeliveryCount: 0,
	}

	var re = regexp.MustCompile(strings.Join(inputs.Keywords, "|"))
	var matchedRecipeNames []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// read opening brace line or first line after a recipe parsed/processed
		scanner.Scan()
		t := scanner.Text()
		// If end of file just break
		if t == "]" {
			break
		}

		postcode := readPostcode(scanner)

		recipeName := readRecipe(scanner)
		_, from, to := readDelivery(scanner)

		if postcode == strconv.Itoa(int(inputs.Postcode)) {
			if inputs.From >= from && inputs.To <= to {
				countPerPostcodeAndTime.DeliveryCount++
			}
		}

		recipeNames[recipeName]++
		postcodes[postcode]++
	}

	// Create and sort CountPerRecipe array
	keys := make([]string, 0, len(recipeNames))
	for k := range recipeNames {
		keys = append(keys, k)
		if re.MatchString(k) {
			matchedRecipeNames = append(matchedRecipeNames, k)
		}
	}
	sort.Sort(sort.StringSlice(keys))
	var cpr []CountPerRecipe
	for _, k := range keys {
		cpr = append(cpr, CountPerRecipe{
			Recipe: k,
			Count:  uint(recipeNames[k]),
		})
	}

	// Find most delivered postcode
	max := 0
	var maxPostCode string
	for p, c := range postcodes {
		if max < c {
			max = c
			maxPostCode = p
		}
	}

	// Match recipes by name

	// Create result
	return &Result{
		UniqueRecipeCount: uint(len(recipeNames)),
		CountPerRecipe:    cpr,
		BusiestPostcode: BusiestPostcode{
			Postcode:      maxPostCode,
			DeliveryCount: uint(max),
		},
		CountPerPostcodeAndTime: countPerPostcodeAndTime,
		MatchByName:             matchedRecipeNames,
	}
}
