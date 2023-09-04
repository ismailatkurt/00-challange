package app

import (
	"encoding/json"
	"testing"
)

func TestCreateResult(t *testing.T) {
	recipe01Name := "Parmesan-Crusted Pork Tenderloin"
	recipe01UniqueCount := uint16(12)
	recipe02Name := "Spanish One-Pan Chicken"
	recipe02UniqueCount := uint16(123)
	recipe03Name := "Speedy Chicken Fajitas"
	recipe03UniqueCount := uint16(1234)

	givenCprs := []CountPerRecipe{
		{
			Recipe: recipe01Name,
			Count:  recipe01UniqueCount,
		},
		{
			Recipe: recipe02Name,
			Count:  recipe02UniqueCount,
		},
		{
			Recipe: recipe03Name,
			Count:  recipe03UniqueCount,
		},
	}

	givenRecipeNames := map[string]uint16{
		recipe01Name: recipe01UniqueCount,
		recipe02Name: recipe02UniqueCount,
		recipe03Name: recipe03UniqueCount,
	}

	countPerPostcodeAndTime := CountPerPostcodeAndTime{
		Postcode:      "10192",
		From:          "10AM",
		To:            "4PM",
		DeliveryCount: 246,
	}
	max := uint32(5432)
	maxPostcode := "13409"
	expectedResult := &Result{
		UniqueRecipeCount: 3,
		CountPerRecipe:    givenCprs,
		BusiestPostcode: BusiestPostcode{
			Postcode:      maxPostcode,
			DeliveryCount: max,
		},
		CountPerPostcodeAndTime: countPerPostcodeAndTime,
		MatchByName: []string{
			"Spanish One-Pan Chicken", "Speedy Chicken Fajitas",
		},
	}

	t.Run("testing Create Result", func(t *testing.T) {
		result := CreateResult(
			givenRecipeNames,
			givenCprs,
			maxPostcode,
			max,
			countPerPostcodeAndTime,
			[]string{
				"Spanish One-Pan Chicken", "Speedy Chicken Fajitas",
			},
		)

		resultJson, err := json.Marshal(result)
		if err != nil {
			return
		}
		expectedResultJson, err := json.Marshal(expectedResult)
		if err != nil {
			return
		}
		if string(resultJson) != string(expectedResultJson) {
			t.Errorf("\ngot %s, \nwant %s\n", resultJson, expectedResultJson)
		}
	})
}
