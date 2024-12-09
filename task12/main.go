package main

import "fmt"

func main() {
	set := make(map[string]bool)

	words := []string{"cat", "cat", "dog", "cat", "tree"}

	for _, val := range words {
		set[val] = true
	}

	for k := range set {
		fmt.Println(k)
	}
}