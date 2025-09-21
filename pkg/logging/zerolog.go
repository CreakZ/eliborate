package logging

import (
	"os"

	"github.com/rs/zerolog"
)

func NewZerologLogger() *zerolog.Logger {
	timeFormat := "15:04:05.000 02.01.2006"

	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: timeFormat,
	}

	logger := zerolog.
		New(output).
		With().
		Timestamp().
		Caller().
		Logger()

	return &logger
}
