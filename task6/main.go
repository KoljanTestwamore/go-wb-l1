package main

import (
	"context"
	"time"
)

func stopGorout1() {
	quit := make(chan bool)

	go func ()  {
		for {
			select {
			case <- quit:
				return
			default:
				// штуки
			}
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		quit<-true
	}()
}

func stopGorout2() {
	ctx, cancel := context.WithCancel(context.Background())

	go func (ctx context.Context)  {
		for {
			select {
			case <-ctx.Done():
					return
			default:
				// приколы
			}

			time.Sleep(500 * time.Millisecond)
	}
	}(ctx)

	go func() {
			time.Sleep(3 * time.Second)
			cancel()
	}()
}

func main() {
	stopGorout1()
	stopGorout2()
}