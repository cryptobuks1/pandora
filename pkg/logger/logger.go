package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"pandora/pkg/cfg"
)

const defaultLevel = zerolog.TraceLevel

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
}

func Initialize() {
	level, err := zerolog.ParseLevel(cfg.Cfg.LogLevel)
	if err != nil {
		log.Warn().Msgf("error while parse log level (%s) setting up default level (%s)", cfg.Cfg.LogLevel, defaultLevel)
		setLogLevel(defaultLevel)
		return
	}
	if level == zerolog.NoLevel {
		log.Warn().Msgf("parsed level is empty so setting up default level (%s)", defaultLevel)
		setLogLevel(defaultLevel)
		return
	}
	log.Debug().Msgf("setting up log level: %s", level)
	setLogLevel(level)
}

func setLogLevel(level zerolog.Level) {
	log.Logger = log.
		Logger.
		Level(level)
	if log.Logger.GetLevel() == zerolog.TraceLevel {
		log.Logger = log.Logger.With().Caller().Logger()
		log.Trace().Msgf("parsed level is %s so added a caller to logger", zerolog.TraceLevel)
	}
}
