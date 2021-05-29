package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}

	return b
}

// func measureResponseTime(url string) time.Duration {
// 	startB := time.Now()
// 	http.Get(url)
// 	return time.Since(startB)

// }
