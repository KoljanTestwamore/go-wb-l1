package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for index := range numbers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("%v\n", numbers[index] * numbers[index])
		}()
	}
	wg.Wait()
}