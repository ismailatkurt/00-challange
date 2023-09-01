package main

import (
	"fmt"
	"os"
)

func openFile() *os.File {
	f, err := os.Open(RecipeFilePath)
	if err != nil {
		_ = fmt.Errorf("could not open the file: %s", err.Error())
	}

	return f
}
