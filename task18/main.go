package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) increment() {
	c.Lock()
	c.count++
	c.Unlock()
}

func main() {
	counter := Counter{}
	wg := sync.WaitGroup{}

	for range 10 {
		wg.Add(1)
		go func ()  {
			counter.increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("%v\n", counter.count)
}