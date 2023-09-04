package app

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

// RawInput used to json unmarshall for given inputs
type RawInput struct {
	Postcode string   `json:"postcode"`
	From     string   `json:"from"`
	To       string   `json:"to"`
	Keywords []string `json:"keywords"`
}

// Input used to compare postcodes and from-to matching. Types added here for better usage.
type Input struct {
	Postcode uint16   `json:"postcode"`
	From     uint8    `json:"from"`
	To       uint8    `json:"to"`
	Keywords []string `json:"keywords"`
}

func getInputs() *Input {
	f := getInputsFile()

	var rawInput RawInput
	err := json.Unmarshal(f, &rawInput)
	if err != nil {
		log.Println(err)
	}

	postcode, err := strconv.ParseUint(rawInput.Postcode, 10, 16)
	if err != nil {
		log.Println(err)
	}
	from, err := strconv.ParseUint(strconv.Itoa(int(convertTo24HourFormat(rawInput.From))), 10, 8)
	if err != nil {
		log.Println(err)
	}
	to, err := strconv.ParseUint(strconv.Itoa(int(convertTo24HourFormat(rawInput.To))), 10, 8)
	if err != nil {
		log.Println(err)
	}

	return &Input{
		Postcode: uint16(postcode),
		From:     uint8(from),
		To:       uint8(to),
		Keywords: rawInput.Keywords,
	}
}

func getInputsFile() []byte {
	f, err := os.ReadFile(InputsFilePath)
	if err != nil {
		_ = fmt.Errorf("could not open inputs file: %s", err.Error())
	}

	return f
}
