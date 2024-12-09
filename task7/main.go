package main

import (
	"fmt"
	"sync"
)

func main() {
	concMap := sync.Map{}
	wg := sync.WaitGroup{}

	for i := range 4 {
		wg.Add(1)
		go func() {
			concMap.Store(i, i + 10)
			wg.Done()
		}()
	}

	wg.Wait()

	concMap.Range(func(key, value any) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})
}