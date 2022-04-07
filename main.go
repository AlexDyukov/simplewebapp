// simplewebapp is a tiny REST api service with /time endpoint
package main

import (
	"fmt"
	"time"
)

func getTime(timezone string) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, fmt.Errorf("time.LoadLocation: %w", err)
	}

	return time.Now().In(loc), nil
}

func main() {
	conf := WebConfig{
		Port: defaultPort,
	}
	conf.ParseParams()

	listenAndServe(conf)
}
