package main

import (
	"context"
	"fmt"
	"time"
)

func Start(workers []func(ctx *context.Context)) {
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(2000 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for _, worker := range workers {
		<-burstyLimiter
		go func(worker func(ctx *context.Context)) {
			fmt.Println("request", time.Now())
			c := context.Background()
			worker(&c)
		}(worker)

	}
}

// 超时处理
func NewWorker(worker func(context.Context, context.CancelFunc)) func(ctx *context.Context) {
	return func(c *context.Context) {
		//c1 := make(chan string, 1)
		//defer close(c1)
		ctx, cancel := context.WithTimeout(*c, 1*time.Second)
		go func() {
			worker(ctx, cancel)
			//c1 <- "worker success"
		}()
	}
}
