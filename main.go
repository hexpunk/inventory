package main

import (
	"io"
	"net/http"
	"net/http/cgi"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var appName = "INVENTORY"

func loadEnvVars() {
	env := strings.ToLower(getAppEnvDefault("ENV", "development"))

	godotenv.Load(".env." + env + ".local")

	if env != "test" {
		godotenv.Load(".env.local")
	}

	godotenv.Load(".env." + env)

	godotenv.Load()
}

func setupLogging() {
	var dest io.Writer
	if isCgiMode() {
		dest = os.Stderr
	} else {
		dest = os.Stdout
	}

	if getAppEnvBool("LOG_JSON") {
		log.Logger = log.Logger.Output(dest)
	} else {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:        dest,
				TimeFormat: time.RFC3339,
				NoColor:    getAppEnvBool("LOG_NO_COLOR"),
			},
		)
	}

	log.Logger.Level(getAppEnvLogLevel())
}

func isCgiMode() bool {
	return getAppEnvBool("CGI_MODE")
}

func main() {
	loadEnvVars()
	setupLogging()

	router := NewRouter()

	if isCgiMode() {
		log.Debug().Msg("Running in CGI mode")
		log.Fatal().Err(
			cgi.Serve(router),
		).Send()
	} else {
		host := getAppEnv("HOST")
		port := getAppEnvDefault("PORT", "8080")

		log.Debug().Str("host", host).Str("port", port).Msg("Running in HTTP mode")
		log.Fatal().Err(
			http.ListenAndServe(host+":"+port, router),
		).Send()
	}
}
