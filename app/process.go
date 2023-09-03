package app

import (
	"regexp"
	"sort"
	"strings"
)

func Process() *Result {
	inputs := getInputs()
	scanner := getFileScanner()

	var re = regexp.MustCompile(strings.Join(inputs.Keywords, "|"))
	var matchedRecipeNames []string

	countPerPostcodeAndTime, recipeNames, postcodes := parseRecipesJson(scanner, *inputs)

	// Create and sort CountPerRecipe array
	keys := make([]string, 0, len(recipeNames))
	for k := range recipeNames {
		keys = append(keys, k)
		if re.MatchString(k) {
			matchedRecipeNames = append(matchedRecipeNames, k)
		}
	}

	// Find most delivered postcode
	max, maxPostCode := getMaxPostCodeAndCount(postcodes)

	// Create result
	return CreateResult(
		recipeNames,
		SortCountPerRecipes(recipeNames, keys),
		maxPostCode,
		max,
		countPerPostcodeAndTime,
		matchedRecipeNames,
	)
}

func SortCountPerRecipes(recipeNames map[string]uint16, keys []string) []CountPerRecipe {
	sort.Sort(sort.StringSlice(keys))
	var cprs []CountPerRecipe
	for _, k := range keys {
		cprs = append(cprs, CountPerRecipe{
			Recipe: k,
			Count:  recipeNames[k],
		})
	}

	return cprs
}
