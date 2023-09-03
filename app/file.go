package app

import (
	"bufio"
	"fmt"
	"os"
)

func getFileScanner() *bufio.Scanner {
	return bufio.NewScanner(openFile())
}

func openFile() *os.File {
	f, err := os.Open(RecipeFilePath)
	if err != nil {
		_ = fmt.Errorf("could not open the file: %s", err.Error())
	}

	return f
}
