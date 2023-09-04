package app

import (
	"encoding/json"
	"fmt"
)

type CountPerRecipe struct {
	Recipe string `json:"recipe"`
	Count  uint16 `json:"count"`
}

type BusiestPostcode struct {
	Postcode      string `json:"postcode"`
	DeliveryCount uint32 `json:"delivery_count"`
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

type IResult interface {
	PrintResults()
}

// PrintResults Just adding indentation
func (r *Result) PrintResults() {
	file, _ := json.MarshalIndent(r, "", " ")

	fmt.Println(string(file))
}

func CreateResult(
	recipeNames map[string]uint16,
	cpr []CountPerRecipe,
	maxPostCode string,
	max uint32,
	countPerPostcodeAndTime CountPerPostcodeAndTime,
	matchedRecipeNames []string,
) *Result {
	return &Result{
		UniqueRecipeCount: uint(len(recipeNames)),
		CountPerRecipe:    cpr,
		BusiestPostcode: BusiestPostcode{
			Postcode:      maxPostCode,
			DeliveryCount: max,
		},
		CountPerPostcodeAndTime: countPerPostcodeAndTime,
		MatchByName:             matchedRecipeNames,
	}
}
