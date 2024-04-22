package main

import (
	"art"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	params := flag.Args()

	if len(params) != 1 || *art.Help {
		fmt.Println("\033[36mUsage: go run . \"[5 #][5 -_]-[5 #]\" will print #####-_-_-_-_-_-#####.\n\nUse -h flag for more info!\033[0m")
		if *art.Help {
			fmt.Println("-m\t\tEnables decoding multiple lines from a file.\n\t\tUse: \033[32mgo run . -m <nameofinputfile>\033[0m")
			fmt.Println("-e\t\tEnables encoding a single line.\n\t\tExample: \033[32mgo run . -e \"#####-_-_-_-_-_-#####\"\033[0m")
			fmt.Println("-me\t\tEnables encoding multiple lines from a file.\n\t\tUse: \033[32mgo run . -me <nameofinputfile>\033[0m")

		}
		return
	}

	input := params[0]
	var content string

	// Check if input resembles a file path
	if _, err := os.Stat(input); err == nil {
		// Input is a file path, read its contents
		absPath, _ := filepath.Abs(input)
		fileBytes, err := os.ReadFile(absPath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		content = string(fileBytes)
	} else {
		// Input is not a file path, check if it's a file in the Resources folder
		absPath, _ := filepath.Abs("Resources/" + input)
		fileBytes, err := os.ReadFile(absPath)
		if err != nil {
			// Input is neither a file path nor a file in the Resources folder, treat it as a string
			content = input
		} else {
			// Input is a file in the Resources folder, read its contents
			content = string(fileBytes)
		}
	}

	// Check for multiline art detection
	if strings.Contains(content, "\n") && !*art.MultilineDecode && !*art.MultilineEncode {
		// Extract file name from the input
		fileName := filepath.Base(input)
		var errMsg string // Declare errMsg outside the if-else blocks
		// Customize error message based on file name
		if strings.Contains(fileName, "encoded") {
			errMsg = art.HandleError("missing_decode_flag").Error()
		} else if strings.Contains(fileName, "art") {
			errMsg = art.HandleError("missing_encode_flag").Error()
		} else {
			errMsg = art.HandleError("missing_flag").Error()
		}
		fmt.Println("\033[31mError:\033[0m", errMsg)
		return
	}

	if *art.MultilineDecode {
		decodedLines, err := art.DecodeMultiLine(content)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(decodedLines)
		return
	} else if *art.SinglelineEncode {
		encodedLine := art.EncodeSingleLine(content)
		fmt.Println(encodedLine)
		return
	} else if *art.MultilineEncode {
		encodedLine := art.EncodeSingleLine(content)
		fmt.Println(encodedLine)
		return
	}

	// If no flag is provided and the input is a single-line string, decode it
	decoded, err := art.DecodeSingleLine(input)
	if err != nil {
		fmt.Println("\033[31mError:\033[0m", err)
	} else {
		fmt.Println(decoded)
	}
}
