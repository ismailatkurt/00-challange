package app

import "testing"

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

	expectedResult := Result{
		UniqueRecipeCount: 3,
		CountPerRecipe:    givenCprs,
		BusiestPostcode: BusiestPostcode{
			Postcode:      "13409",
			DeliveryCount: 5342,
		},
		CountPerPostcodeAndTime: CountPerPostcodeAndTime{
			Postcode:      "10192",
			From:          "10AM",
			To:            "4PM",
			DeliveryCount: 246,
		},
		MatchByName: []string{
			"Spanish One-Pan Chicken", "Speedy Chicken Fajitas",
		},
	}

	result := CreateResult(
		givenRecipeNames,
		givenCprs,
		"123",
		3,
		CountPerPostcodeAndTime{
			Postcode:      "10192",
			From:          "10AM",
			To:            "4PM",
			DeliveryCount: 246,
		},
		[]string{
			"Spanish One-Pan Chicken", "Speedy Chicken Fajitas",
		},
	)

}
