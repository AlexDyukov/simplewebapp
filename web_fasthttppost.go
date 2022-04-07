//go:build fasthttppost

package main

import (
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func getTimezone(ctx *fasthttp.RequestCtx) (string, error) {
	timezone := string(ctx.PostBody())
	if len(timezone) == 0 {
		timezone = "UTC"
	}

	return timezone, nil
}

func serve(ctx *fasthttp.RequestCtx) {
	timezone, err := getTimezone(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	now, err := getTime(timezone)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	fmt.Fprintln(ctx, now)
}

func listenAndServe(conf WebConfig) {
	r := router.New()

	r.POST("/time", serve)

	server := &fasthttp.Server{
		Handler: r.Handler,
	}

	if err := server.ListenAndServe(fmt.Sprintf(":%d", conf.Port)); err != nil {
		log.Fatal(err)
	}
}
