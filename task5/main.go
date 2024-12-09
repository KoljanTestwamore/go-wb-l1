package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// количество секунд
	N := 5

	c := make(chan int)

	mut := sync.Mutex{}

	go func() {
		for {
			j, ok := <-c
			if (!ok) {
				return
			}
			fmt.Printf("Recieved %v\n", j)
		}
	}()

	mut.Lock()
	go func() {
		var i int

		for {
			i++
			c<-i
			fmt.Printf("Sent %v\n", i)
			time.Sleep(time.Second)

			if (i >= N) {
				mut.Unlock()
				return
			}
		}
	}()

	mut.Lock()
	close(c)
}
