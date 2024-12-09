package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	var result int
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	for index := range numbers {
		wg.Add(1)
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			defer wg.Done()

			result += numbers[index] * numbers[index]
		}()
	}
	wg.Wait()

	fmt.Printf("%d", result)
}