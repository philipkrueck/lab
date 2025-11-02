package racer

import (
	"fmt"
	"net/http"
	"time"
)

var timeLimit = time.Duration(10 * time.Second)

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, timeLimit)
}

func ConfigurableRacer(a, b string, timeLimit time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		winner = a
	case <-ping(b):
		winner = b
	case <-time.After(timeLimit):
		err = fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

	return
}

func ping(url string) (ch chan bool) {
	ch = make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
