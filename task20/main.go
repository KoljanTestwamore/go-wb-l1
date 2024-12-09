package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "snow dog sun"

	words := strings.Fields(input)

	for i := range words {
		fmt.Printf("%v ", string(words[len(words) - i - 1]))
	}
}
