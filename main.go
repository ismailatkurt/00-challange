package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const RecipeFilePath = "recipes.json"
const InputsFilePath = "inputs.json"
const OutputFilePath = "output.json"

func main() {
	t1 := time.Now()
	result := Process()

	file, _ := json.MarshalIndent(result, "", " ")

	_ = os.WriteFile(OutputFilePath, file, 0644)
	t2 := time.Now()

	diff := t2.Sub(t1)
	fmt.Println(diff)
}
