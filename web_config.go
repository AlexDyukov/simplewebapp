package main

import (
	"flag"
	"log"
)

type WebConfig struct {
	Port int
}

const (
	defaultPort = 8080
)

func (cfg *WebConfig) ParseParams() {
	flag.IntVar(&cfg.Port, "port", defaultPort, "http listen port, 1025-65535")
	flag.Parse()

	if cfg.Port < 1025 || cfg.Port > 65535 {
		fmtStr := "invalid value \"%d\" for flag -port: should be in range [1025;65535]\n"
		log.Fatalf(fmtStr, cfg.Port)
	}
}
