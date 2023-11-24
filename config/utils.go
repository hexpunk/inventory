package config

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

func getEnv(appName, key string) string {
	return strings.ToLower(strings.TrimSpace(os.Getenv(appName + "_" + key)))
}

func orDefault(value, def string) string {
	if value == "" {
		return def
	}

	return value
}

func toBool(value string) bool {
	value = strings.ToLower(strings.TrimSpace(value))
	return value != "false" && value != ""
}

func toLogLevel(level string) zerolog.Level {
	switch level {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	case "disabled":
		return zerolog.Disabled
	default:
		return zerolog.InfoLevel
	}
}
