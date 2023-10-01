package _select

import (
	"net/http"
	"time"
)

func Racer(url1 string, url2 string) (winner string) {

	aDuration := measureResponseTime(url1)
	bDuration := measureResponseTime(url2)

	if aDuration < bDuration {
		return url1
	}
	return url2
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
