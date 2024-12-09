package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	fmt.Print("Enter workers' amount: ")
	var workerNumber int
	fmt.Scanf("%d\n", &workerNumber)

	wg := sync.WaitGroup{}

	c := make(chan int)

	for index := range workerNumber {
		go func() {
			fmt.Printf("Starting worker %v...\n", index)
			wg.Add(1)
				
			for {
				j, ok := <-c
				if (!ok) {
					wg.Done()
					fmt.Printf("Worker %v finished\n", index)

					return
				}

				fmt.Printf("Worker %v, Value %v\n", index, j)
			}
		}()
	}

	go func() {
		var i int

		for {
			i++
			c<-i
			time.Sleep(time.Second)
		}
	}()

	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT)
	
	<-termChan // Blocks here until interrupted
	
	fmt.Printf("****\nShutdown signal recieved\n****")
	close(c)
	wg.Wait()
}
