package main_test

import (
	"testing"
	"time"
)

func TestLoadLocationConcurrency(t *testing.T) {
	const (
		testCount    = 1024
		testLocation = "Asia/Tokyo"
	)

	ch := make(chan struct{})
	defer close(ch)

	for i := 0; i < testCount; i++ {
		go func() {
			for i := 0; i < testCount; i++ {
				_, err := time.LoadLocation(testLocation)
				if err != nil {
					panic("cannot load test timezone " + err.Error())
				}
			}
			ch <- struct{}{}
		}()
	}

	for i := 0; i < testCount; i++ {
		<-ch
	}
}
