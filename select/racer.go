package main

import (
	"net/http"
)

func Racer(url_a, url_b string) (winner string) {
	// select waits on multiple channels
	// here first to return is returned by the select
	select {
	case <-ping(url_a):
		return url_a
	case <-ping(url_b):
		return url_b
	}
}

func ping(url string) chan struct{} {
	// Always use `make` when creating a channel as just initalising a variable
	// will give the "zero" value of the type
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
