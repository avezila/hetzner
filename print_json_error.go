package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func PrintJSONError(data []byte, err error) {
	if jsonError, ok := err.(*json.SyntaxError); ok {
		line, character, lcErr := lineAndCharacter(data, int(jsonError.Offset))
		fmt.Fprintf(os.Stderr, "json syntax error at line %d, character %d: %v\n", line, character, jsonError.Error())
		if lcErr != nil {
			fmt.Fprintf(os.Stderr, "Couldn't find the line and character position of the error due to error %v\n", lcErr)
		}
		return
	}
	if jsonError, ok := err.(*json.UnmarshalTypeError); ok {
		line, character, lcErr := lineAndCharacter(data, int(jsonError.Offset))
		fmt.Fprintf(os.Stderr, "The JSON type '%v' cannot be converted into the Go '%v' type on struct '%s', field '%v'. See input file line %d, character %d\n", jsonError.Value, jsonError.Type.Name(), jsonError.Struct, jsonError.Field, line, character)
		if lcErr != nil {
			fmt.Fprintf(os.Stderr, "Couldn't find the line and character position of the error due to error %v\n", lcErr)
		}
		return
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "json error: %v\n", err)
		return
	}
}

func lineAndCharacter(binput []byte, offset int) (line int, character int, err error) {
	input := string(binput)
	lf := rune(0x0A)

	if offset > len(input) || offset < 0 {
		return 0, 0, fmt.Errorf("Couldn't find offset %d within the input.", offset)
	}

	// Humans tend to count from 1.
	line = 1

	for i, b := range input {
		if b == lf {
			line++
			character = 0
		}
		character++
		if i == offset {
			break
		}
	}

	return line, character, nil
}
