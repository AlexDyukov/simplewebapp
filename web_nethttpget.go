//go:build !nethttppost

package main

import (
	"fmt"
	"net/http"

	"github.com/alexdyukov/currtime"
	"github.com/julienschmidt/httprouter"
)

func getTimezone(r *http.Request) string {
	query := r.URL.Query()
	timezone := query.Get("timezone")

	return timezone
}

func getTime(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	timezone := getTimezone(r)

	currentTime, err := currtime.GetTime(timezone)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, currentTime)
}

func listenAndServe(conf WebConfig) {
	router := httprouter.New()

	router.GET("/time", getTime)

	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), router)
}
