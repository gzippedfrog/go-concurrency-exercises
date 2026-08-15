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
	jobs := make(chan struct{}, numRequests)

	for range numWorkers {
		wg.Go(func() {
			for range jobs {
				networkRequest()
			}
		})
	}

	for range numRequests {
		jobs <- struct{}{}
	}

	close(jobs)

	wg.Wait()

	fmt.Println(atomic.LoadInt64(&count))
}
