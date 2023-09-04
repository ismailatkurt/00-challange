package app

import (
	"bufio"
	"strings"
)

type Reader struct {
	s *bufio.Scanner
}

type IReader interface {
	ReadPostcode() string
	ReadRecipe() string
	ReadDelivery() (string, uint8, uint8)
	getPostcodeFromLine(l *string) string
	getRecipeFromLine(l *string) string
	getDeliveryFromLine(l *string) string
}

func (r *Reader) ReadPostcode() string {
	r.s.Scan()
	postcodeLine := r.s.Text()

	return r.getPostcodeFromLine(&postcodeLine)
}

func (r *Reader) ReadRecipe() string {
	r.s.Scan()
	recipeLine := r.s.Text()

	return r.getRecipeFromLine(&recipeLine)
}

func (r *Reader) ReadDelivery() (string, uint8, uint8) {
	r.s.Scan()
	deliveryLine := r.s.Text()
	deliveryLine = r.getDeliveryFromLine(&deliveryLine)

	// Here and the next line, I use strings.Cut and .Split
	// but probably without them, I could have saved 20% of execution time
	day, timeInterval, _ := strings.Cut(deliveryLine, " ")

	parts := strings.Split(timeInterval, " - ")
	fromStr, toStr := parts[0], parts[1][:len(parts[1])]

	from := convertTo24HourFormat(fromStr)
	to := convertTo24HourFormat(toStr)

	return day, from, to
}

func (r *Reader) getPostcodeFromLine(l *string) string {
	postcode := strings.Trim(*l, "\t\n\v\f\r ")
	postcode = postcode[13:]

	// Initially I have used strings.TrimSuffix but it has so bad performance.
	// Then I assumed the structure is always in correct way and decided to use this approach
	return postcode[:len(postcode)-2]
}

func (r *Reader) getRecipeFromLine(l *string) string {
	recipe := strings.Trim(*l, "\t\n\v\f\r ")
	recipe = recipe[11:]

	return recipe[:len(recipe)-2]
}

func (r *Reader) getDeliveryFromLine(l *string) string {
	delivery := strings.Trim(*l, "\t\n\v\f\r ")
	delivery = delivery[13:]

	return delivery[:len(delivery)-1]
}
