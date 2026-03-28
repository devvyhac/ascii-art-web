package main

import (
	"fmt"
	"os"
	"strings"
)

func Art(str, font string) (string, error) {

	words := strings.Split(str, "\\n")
	var builder strings.Builder

	file, err := os.ReadFile(font)
	if err != nil {
		return "", err
	}

	content := strings.Split(string(file), "\n")
	if font == "thinkertoy.txt" {
		content = strings.Split(string(file), "\r\n")
	}

	for i, word := range words {
		if words[len(words)-1] == "" && i == len(words)-1 && words[i-1] == "" {
			break
		}

		if word == "" {
			fmt.Println()
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range word {
				builder.WriteString(content[FetchIndex(char)+i])
			}
			builder.WriteString("\n")
		}
	}
	result := builder.String()
	return result, nil
	// Reverse()
}

func FetchIndex(char rune) int {
	return 9*(int(char)-32) + 1
}
