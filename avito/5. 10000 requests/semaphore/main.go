package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const numRequests = 10000
const numWorkers = 100

var count int64

func networkRequest() {
	time.Sleep(time.Millisecond)
	atomic.AddInt64(&count, 1)
}

func main() {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, numWorkers)

	for range numRequests {
		wg.Go(func() {
			semaphore <- struct{}{}
			defer func() { <-semaphore }()
			networkRequest()
		})
	}

	wg.Wait()

	fmt.Println(count)
}
