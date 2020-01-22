package main

import (
	"net/http"
	"time"
)

func Racer(url_a, url_b string) (winner string) {
	if duration(url_a) < duration(url_b) {
		return url_a
	}

	return url_b
}

func duration(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
