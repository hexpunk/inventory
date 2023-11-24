package main

import (
	"io"
	"net/http"
	"net/http/cgi"
	"os"
	"time"

	"github.com/hexpunk/inventory/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var appName = "INVENTORY"

func setupLogging(config *config.Config) {
	var dest io.Writer
	if config.CgiMode() {
		dest = os.Stderr
	} else {
		dest = os.Stdout
	}

	if config.LogJson() {
		log.Logger = log.Logger.Output(dest)
	} else {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:        dest,
				TimeFormat: time.RFC3339,
				NoColor:    config.LogNoColor(),
			},
		)
	}

	log.Logger.Level(config.LogLevel())
}

func main() {
	config := config.GetConfig(appName)
	setupLogging(config)

	router := NewRouter()

	if config.CgiMode() {
		log.Debug().Msg("Running in CGI mode")
		log.Fatal().Err(
			cgi.Serve(router),
		).Send()
	} else {
		host := config.Host()
		port := config.Port()

		log.Debug().Str("host", host).Str("port", port).Msg("Running in HTTP mode")
		log.Fatal().Err(
			http.ListenAndServe(host+":"+port, router),
		).Send()
	}
}
