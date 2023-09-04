package app

import (
	"regexp"
	"sort"
	"strings"
)

func Process() *Result {
	inputs := getInputs()
	scanner := getFileScanner()

	// I think regexp has better performance than strings.Contains
	var re = regexp.MustCompile(strings.Join(inputs.Keywords, "|"))
	var matchedRecipeNames []string

	countPerPostcodeAndTime, recipeNames, postcodes := parseRecipesJson(scanner, *inputs)

	// Create and sort CountPerRecipe array
	// That could have been move to its own function, but it is fine for "not production" version
	keys := make([]string, 0, len(recipeNames))
	for k := range recipeNames {
		keys = append(keys, k)
		// I wanted to take care of Matched Recipe Names here instead of in a different loop/function for the sake of performance
		if re.MatchString(k) {
			matchedRecipeNames = append(matchedRecipeNames, k)
		}
	}

	// Find most delivered postcode
	max, maxPostCode := getMaxPostCodeAndCount(postcodes)

	// Create and return result
	return CreateResult(
		recipeNames,
		sortCountPerRecipes(recipeNames, keys),
		maxPostCode,
		max,
		countPerPostcodeAndTime,
		matchedRecipeNames,
	)
}

// Maybe I should move this to result.go file
func sortCountPerRecipes(recipeNames map[string]uint16, keys []string) []CountPerRecipe {
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
