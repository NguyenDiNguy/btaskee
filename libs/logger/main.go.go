package logger

import (
	"os"

	"github.com/UnnoTed/horizontal"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	log.Logger = log.Output(horizontal.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()
	log.Level(zerolog.TraceLevel)
}
