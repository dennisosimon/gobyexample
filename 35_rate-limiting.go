package main

import (
	"time"
	"fmt"
)

func main() {
	request := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		request <- i
	}
	close(request)

	limiter := time.Tick(1000 * time.Millisecond)
	for req := range request {
		<- limiter
		fmt.Println("request", req, time.Now())
	}


	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(1000 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	bustryRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		bustryRequest <- i
	}
	close(bustryRequest)

	for req := range bustryRequest {
		<- burstyLimiter
		fmt.Println("bustryRequest", req, time.Now())
	}
}
