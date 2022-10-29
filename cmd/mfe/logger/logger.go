package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var App zerolog.Logger

func init() {
	// set default logger configurations
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	// set the default logger level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// add filename to the default logger
	log.Logger = log.With().Caller().Logger()

	App = log.With().Str("type", "Application").Logger()

}
