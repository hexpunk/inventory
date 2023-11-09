package main

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

func getAppEnv(key string) string {
	return os.Getenv(appName + "_" + key)
}

func getAppEnvDefault(key, def string) string {
	env := getAppEnv(key)
	if env == "" {
		return def
	}

	return env
}

func getAppEnvBool(key string) bool {
	value := strings.ToLower(strings.TrimSpace(getAppEnv(key)))
	return value != "false" && value != ""
}

func getAppEnvLogLevel() zerolog.Level {
	switch level := strings.ToLower(strings.TrimSpace(getAppEnv("LOG_LEVEL"))); level {
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
