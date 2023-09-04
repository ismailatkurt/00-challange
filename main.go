package main

import (
	"example.com/recipe-app/app"
	"time"
)

func main() {
	// to be used to calculate execution time
	t1 := time.Now()

	result := app.Process()
	result.PrintResults()

	t2 := time.Now()
	diff := t2.Sub(t1)
	app.PrintExecutionDetails(diff)
}
