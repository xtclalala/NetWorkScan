package main

import (
	"context"
	"sync"
)

func Start(workers []func()) {
	burst := global.Burst.BurstNum
	burstLimiter := make(chan func())
	defer close(burstLimiter)
	c, cancel := context.WithCancel(context.Background())

	go func() {
		for i := 0; i < len(workers); i++ {
			burstLimiter <- workers[i]
		}
		cancel()
	}()

	wg := &sync.WaitGroup{}
	for i := 0; i < burst; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case w := <-burstLimiter:
					w()
				case <-c.Done():
					return
				}
			}
		}()
	}
	wg.Wait()
}
