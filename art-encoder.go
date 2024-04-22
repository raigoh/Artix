package art

import (
	"fmt"
	"strings"
)

// EncodeSingleLine encodes a single string line
func EncodeSingleLine(line string) string {
	// Initialize variables
	lineLen := len(line)
	var changeList []string
	var startIndices []int
	var endIndices []int

	// Find consecutive characters and store them in changeList
	counter := 0
	for counter < lineLen-1 {
		endCounter := findEndOfConsecutiveChars(line, counter)
		if counter != endCounter {
			changeList = append(changeList, line[counter:endCounter+1])
			startIndices = append(startIndices, counter)
			endIndices = append(endIndices, endCounter)
		}
		counter = endCounter + 1
	}

	// Replace consecutive characters with encoded format
	for i := len(changeList) - 1; i >= 0; i-- {
		charCount := len(changeList[i])
		encodedFormat := fmt.Sprintf("[%d %s]", charCount, string(changeList[i][0]))
		line = replaceSubstring(line, encodedFormat, startIndices[i], endIndices[i])
	}

	// Find consecutive pairs and store them in changeList
	pairCounter := 0
	var pairList []string
	var pairStartIndices []int
	var pairEndIndices []int

	for pairCounter < lineLen-4 {
		pairEndCounter := findEndOfConsecutivePairs(line, pairCounter) + 1
		if pairCounter+1 != pairEndCounter {
			pairList = append(pairList, line[pairCounter:pairEndCounter+1])
			pairStartIndices = append(pairStartIndices, pairCounter)
			pairEndIndices = append(pairEndIndices, pairEndCounter)
		}
		pairCounter = pairEndCounter
	}

	// Replace consecutive pairs with encoded format
	for i := len(pairList) - 1; i >= 0; i-- {
		charCount := len(pairList[i])
		encodedFormat := fmt.Sprintf("[%d %s%s]", charCount/2, string(pairList[i][0]), string(pairList[i][1]))
		line = replaceSubstring(line, encodedFormat, pairStartIndices[i], pairEndIndices[i])
	}

	return line
}

// findEndOfConsecutiveChars finds the end index of consecutive characters starting from the given index
func findEndOfConsecutiveChars(line string, startIndex int) int {
	endIndex := startIndex
	for endIndex < len(line)-1 && line[startIndex] == line[endIndex+1] {
		endIndex++
	}
	return endIndex
}

// findEndOfConsecutivePairs finds the end index of consecutive pairs starting from the given index
func findEndOfConsecutivePairs(line string, startIndex int) int {
	endIndex := startIndex
	for endIndex < len(line)-4 && line[startIndex:startIndex+2] == line[endIndex+2:endIndex+4] {
		endIndex += 2
	}
	return endIndex
}

// replaceSubstring replaces the substring in the given range with the new substring
func replaceSubstring(line string, newSubstr string, startIndex int, endIndex int) string {
	return line[:startIndex] + newSubstr + line[endIndex+1:]
}

// EncodeMultiLine function encodes multiple string lines
func EncodeMultiLine(content string) (string, error) {
	lines := strings.Split(content, "\n")
	var encodedLines []string
	for _, line := range lines {
		encodedLine := EncodeSingleLine(line)
		encodedLines = append(encodedLines, encodedLine)
	}
	return strings.Join(encodedLines, "\n"), nil
}
