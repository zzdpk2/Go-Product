package _select

import (
	"net/http"
	"time"
)

func Racer(url1 string, url2 string) (winner string) {

	// aDuration := measureResponseTime(url1)
	// bDuration := measureResponseTime(url2)
	//
	// if aDuration < bDuration {
	// 	return url1
	// }
	// return url2

	select {
	case <-ping(url1):
		return url1
	case <-ping(url2):
		return url2
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
