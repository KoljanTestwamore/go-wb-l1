package main

import "fmt"

func main() {
	input := []rune("главрыба")

	for i := range input {
		fmt.Printf("%v", string(input[len(input) - i - 1]))
	}
}