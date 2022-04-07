//go:build !nethttppost && !fasthttpget && !fasthttppost

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

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

func getTimezone(request *http.Request) (string, error) {
	timezone := request.URL.Query().Get("timezone")
	if timezone == "" {
		timezone = "UTC"
	}

	return timezone, nil
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

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
