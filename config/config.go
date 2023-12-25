package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

// Immutable
type Config struct {
	environment string

	// Server-related settings

	cgiMode bool

	host string
	port string

	// Database-related settings

	sqliteDsn string

	// Log-related settings

	logLevel   zerolog.Level
	logJson    bool
	logNoColor bool
}

func (c *Config) Environment() string {
	return c.environment
}

func (c *Config) CgiMode() bool {
	return c.cgiMode
}

func (c *Config) Host() string {
	return c.host
}

func (c *Config) Port() string {
	return c.port
}

func (c *Config) SqliteDsn() string {
	return c.sqliteDsn
}

func (c *Config) LogLevel() zerolog.Level {
	return c.logLevel
}

func (c *Config) LogJson() bool {
	return c.logJson
}

func (c *Config) LogNoColor() bool {
	return c.logNoColor
}

func GetConfig(appName string) *Config {
	env := orDefault(getEnv(appName, "ENVIRONMENT"), "development")

	godotenv.Load(".env." + env + ".local")

	// TODO: Should this be "test" or "testing"? Does Go's test runner set anything?
	if env != "test" {
		godotenv.Load(".env.local")
	}

	godotenv.Load(".env." + env)

	godotenv.Load()

	return &Config{
		environment: env,

		cgiMode: toBool(getEnv(appName, "CGI_MODE")),

		host: getEnv(appName, "HOST"),
		port: orDefault(getEnv(appName, "PORT"), "8080"),

		sqliteDsn: orDefault(getEnv(appName, "SQLITE_DSN"), ":memory:"),

		logLevel:   toLogLevel(getEnv(appName, "ENVIRONMENT")),
		logJson:    toBool(getEnv(appName, "LOG_JSON")),
		logNoColor: toBool(getEnv(appName, "LOG_NO_COLOR")),
	}
}
