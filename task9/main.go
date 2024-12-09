package main

import (
	"fmt"
	"sync"
)

func main() {
	vals := []int{5, 2, 9}

	c := make(chan int)
	c2 := make(chan int)

	// для одной горутины можно и просто мютекс
	mut := sync.Mutex{}

	go func() {
		for {
			val := <- c
			c2 <- val*val
		}
	}()

	go func() {
		mut.Lock()
		for {
			val, ok := <- c2

			if (!ok) {
				mut.Unlock()
				return
			}

			fmt.Println(val)
		}
	}()

	for _, el := range vals {
		c <- el
	}
	
	close(c2)
	mut.Lock()
}