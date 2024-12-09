package main

import (
	"fmt"
)

func main() {
	tmpMap := make(map[int][]float64)

	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, el := range temperatures {
		index := int(el / 10) * 10

		tmpMap[index] = append(tmpMap[index], el)
	}

	for key, el := range tmpMap {
		fmt.Printf("%v: {%v}\n", key, el)
	}
}