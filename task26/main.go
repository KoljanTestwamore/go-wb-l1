package main

import (
	"fmt"
	"unicode"
)

func main() {
	dict := make(map[rune]bool)

	var input string

	fmt.Scanf("%s\n", &input)

	for _, c := range input {
		char := unicode.ToLower(c)

		_, ok := dict[char]
		if (!ok) {
			dict[char] = true
			continue
		}

		fmt.Printf("false")
		return
	}

	fmt.Printf("true")
}