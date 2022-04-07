//go:build nethttppost

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getTimezone(request *http.Request) (string, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return "", err
	}

	timezone := string(body)
	if len(timezone) == 0 {
		timezone = "UTC"
	}

	return timezone, nil
}

type handler struct{}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	timezone, err := getTimezone(request)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	now, err := getTime(timezone)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprintln(responseWriter, now)
}

func listenAndServe(conf WebConfig) {
	mux := http.NewServeMux()
	mux.Handle("/time", &handler{})

	timeout := time.Second

	server := &http.Server{
		Addr:                         fmt.Sprintf(":%d", conf.Port),
		Handler:                      mux,
		TLSConfig:                    nil,
		ReadTimeout:                  timeout,
		ReadHeaderTimeout:            timeout,
		WriteTimeout:                 timeout,
		IdleTimeout:                  timeout,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
		DisableGeneralOptionsHandler: false,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
