//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream, tweets chan *Tweet, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		tweet, err := stream.Next()

		if err == ErrEOF {
			quit <- true
			return
		}

		tweets <- tweet
	}
}

func consumer(tweets chan *Tweet, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case t := <-tweets:
			if t.IsTalkingAboutGo() {
				fmt.Println(t.Username, "\ttweets about golang")
			} else {
				fmt.Println(t.Username, "\tdoes not tweet about golang")
			}
		case <-quit:
			return
		}
	}
}

func main() {
	start := time.Now()
	tweets := make(chan *Tweet)
	quit := make(chan bool)
	var wg sync.WaitGroup

	// Producer
	wg.Add(1)
	go producer(GetMockStream(), tweets, quit, &wg)

	// Consumer
	wg.Add(1)
	go consumer(tweets, quit, &wg)
	wg.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))
}
