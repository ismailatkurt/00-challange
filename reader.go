package main

import (
	"bufio"
	"strings"
)

func readPostcode(s *bufio.Scanner) string {
	s.Scan()
	postcodeLine := s.Text()

	return getPostcodeFromLine(&postcodeLine)
}

func readRecipe(s *bufio.Scanner) string {
	s.Scan()
	recipeLine := s.Text()

	return getRecipeFromLine(&recipeLine)
}

func readDelivery(s *bufio.Scanner) (string, uint8, uint8) {
	s.Scan()
	deliveryLine := s.Text()
	deliveryLine = getDeliveryFromLine(&deliveryLine)

	day, timeInterval, _ := strings.Cut(deliveryLine, " ")

	parts := strings.Split(timeInterval, " - ")
	//fromStr, toStr := parts[0], strings.TrimSuffix(parts[1], "\"")
	fromStr, toStr := parts[0], parts[1][:len(parts[1])-1]

	from := ConvertTo24HourFormat(fromStr)
	to := ConvertTo24HourFormat(toStr)

	return day, from, to
}

func getPostcodeFromLine(l *string) string {
	postcode := strings.Trim(*l, "\t\n\v\f\r ")
	postcode = postcode[13:]

	return strings.TrimSuffix(postcode, "\",")
}

func getRecipeFromLine(l *string) string {
	recipe := strings.Trim(*l, "\t\n\v\f\r ")
	recipe = recipe[11:]

	return strings.TrimSuffix(recipe, "\",")
}

func getDeliveryFromLine(l *string) string {
	delivery := strings.Trim(*l, "\t\n\v\f\r ")
	delivery = delivery[13:]

	//return strings.TrimSuffix(delivery, "\",")
	return delivery[:len(delivery)-1]
}
