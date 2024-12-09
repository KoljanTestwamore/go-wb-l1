package main

import "fmt"

func delete(slice *[]int, i int) {
	val := *slice

	*slice = append(val[:i], val[i+1:]...)
}

func main() {
	elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	delete(&elements, 3)

	fmt.Printf("%v", elements)
}