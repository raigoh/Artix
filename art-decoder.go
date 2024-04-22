package art

import (
	"strconv"
	"strings"
)

// DecodeSingleLine function decodes a single string line
func DecodeSingleLine(line string) (string, error) {
	// Check if the number of opening and closing brackets is balanced
	openingCount := strings.Count(line, "[")
	closingCount := strings.Count(line, "]")
	if openingCount != closingCount {
		return "", HandleError("un_brackets")
	}

	decodedLine := line

	// Iterate through each bracket pair in the line
	for strings.Contains(decodedLine, "[") && strings.Contains(decodedLine, "]") {
		openingBracketIndex := strings.Index(decodedLine, "[")
		closingBracketIndex := strings.Index(decodedLine, "]")

		// Extract the substring between the brackets
		bracketedSubstring := decodedLine[openingBracketIndex+1 : closingBracketIndex]

		// Process the arguments inside the brackets
		result, err := bracketChange(bracketedSubstring)
		if err != nil {
			return "", err
		}

		// Replace the original bracketed substring with the result
		decodedLine = decodedLine[:openingBracketIndex] + result + decodedLine[closingBracketIndex+1:]
	}

	return decodedLine, nil
}

func bracketChange(arguments string) (string, error) {
	// Check for balanced brackets
	if strings.Count(arguments, "[") != strings.Count(arguments, "]") {
		return "", HandleError("un_brackets")
	}

	// Check for nested brackets
	if strings.Index(arguments, "[") < strings.Index(arguments, "]") {
		return "", HandleError("un_brackets")
	}

	// Find the index of the first space character
	spaceIndex := strings.Index(arguments, " ")

	// Check if space exists
	if spaceIndex == -1 {
		return "", HandleError("no_space")
	}

	// Check if the argument is not empty
	if spaceIndex == len(arguments)-1 {
		return "", HandleError("missing_argument")
	}

	// Split the arguments into number and content
	numberStr := arguments[:spaceIndex]
	content := arguments[spaceIndex+1:]

	// Convert the number string to an integer
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return "", HandleError("not_a_number")
	}

	// Repeat the content number times
	result := strings.Repeat(content, number)

	return result, nil
}

// DecodeMultiLine function decodes multiple string lines
func DecodeMultiLine(content string) (string, error) {
	lines := strings.Split(content, "\n")
	var decodedLines []string
	for _, line := range lines {
		decodedLine, err := DecodeSingleLine(line)
		if err != nil {
			return "", err
		}
		decodedLines = append(decodedLines, decodedLine)
	}
	return strings.Join(decodedLines, "\n"), nil
}
