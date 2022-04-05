package main

import (
	"fmt"
	"net/http"

	"github.com/alexdyukov/currtime"
	"github.com/julienschmidt/httprouter"
)

func getTime(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query()
	timezone := query.Get("timezone")

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
