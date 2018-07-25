package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/hashicorp/consul/api"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
)

var (
	logger              zerolog.Logger
	lambdaName          string
	lambdaID            xid.ID
	consulUrl           string
	version             string
	consulApiClient     *api.Client
	httpClient          *http.Client
	consulConfiguration Config
	initialized         = false
)

func initialize() error {
	lambdaID = xid.New()
	lambdaName = os.Getenv("AWS_LAMBDA_FUNCTION_NAME")
	zerolog.TimestampFieldName = "timestamp"
	logger = zerolog.New(os.Stdout).With().
		Timestamp().
		Str("lambdaID", lambdaID.String()).
		Str("lambdaName", lambdaName).
		Str("environment", os.Getenv("ENVIRONMENT")).
		Str("version", version).
		Logger()

	levelStr := os.Getenv("LOG_LEVEL")
	if levelStr == "" {
		levelStr = "DEBUG"
	}
	level, err := stringToLogLevel(levelStr)

	if err != nil {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		logger.Error().
			Str("input", levelStr).
			Msg("Invalid logging level specified - logging at debug level")
	} else {
		zerolog.SetGlobalLevel(level)
	}

	logger.Info().Msg(fmt.Sprintf("Initializing lambda %s with ID: %s", lambdaName, lambdaID.String()))

	// Pass in the consul url for the environment here. This should be configured in terraform lambda function
	consulUrl = os.Getenv("CONSUL_URL")

	// Put any consul variables needed to be used in this struct. Follow the
	// convention used here for the annotation.
	config := &struct {
		VaultUrl string `consul:"vault/url"`
	}{}

	initialized = true
	return nil
}

func stringToLogLevel(logLevel string) (zerolog.Level, error) {
	switch logLevel {
	case "DEBUG":
		return zerolog.DebugLevel, nil
	case "INFO":
		return zerolog.InfoLevel, nil
	case "WARN":
		return zerolog.WarnLevel, nil
	case "ERROR":
		return zerolog.ErrorLevel, nil
	case "FATAL":
		return zerolog.FatalLevel, nil
	case "PANIC":
		return zerolog.PanicLevel, nil
	case "DISABLED":
		return zerolog.Disabled, nil
	default:
		return zerolog.DebugLevel, errors.New("invalid log level")
	}
}
