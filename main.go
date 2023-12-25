package main

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/cgi"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/hexpunk/inventory/config"
	"github.com/hexpunk/inventory/db"
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

func setupDatabase(config *config.Config) (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", config.SqliteDsn())
	if err != nil {
		return nil, err
	}

	if strings.Contains(config.SqliteDsn(), ":memory:") {
		log.Warn().Msg("In-memory SQLite database detected. Data will not be persisted!")
	}

	// Enforce foreign keys by default
	if _, err := conn.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		return nil, err
	}

	if err = db.Migrate(conn, migrations); err != nil {
		return nil, err
	}

	return conn, nil
}

func main() {
	config := config.GetConfig(appName)

	setupLogging(config)

	conn, err := setupDatabase(config)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer conn.Close()

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
