//go:build nethttppost

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alexdyukov/currtime"
	"github.com/julienschmidt/httprouter"
)

func getTimezone(r *http.Request) (string, error) {
	var tz struct {
		Timezone string `json:"timezone"`
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(body, &tz); err != nil {
		return "", err
	}

	return tz.Timezone, nil
}

func getTime(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	timezone, err := getTimezone(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	currentTime, err := currtime.GetTime(timezone)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, currentTime)
}

func listenAndServe(conf WebConfig) {
	router := httprouter.New()

	router.POST("/time", getTime)

	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), router)
}
