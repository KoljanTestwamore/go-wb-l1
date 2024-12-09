package main

import (
	"fmt"
	"time"
)

func sleep(t time.Duration) {
	start := time.Now()

	for {
		if time.Since(start) > t {
			return
		}
	}
}

func main() {
	fmt.Println("HH:mm:ss", time.Now())
	sleep(5*time.Second)
	fmt.Println("HH:mm:ss", time.Now())
}