package main

import (
    "fmt"
    "time"
)

const numRequests = 10000

var count int

func networkRequest() {
    time.Sleep(time.Millisecond)
    count++
}

func main() {
    for i := 0; i < numRequests; i++ {
        networkRequest()
    }
    fmt.Println(count)
}