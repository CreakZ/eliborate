package log

import (
	"errors"
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

// add zerolog instead
type Log struct {
	InfoLogger  *zerolog.Logger
	ErrorLogger *zerolog.Logger
}

func InitLogger() (*Log, *os.File, *os.File) {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"

	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		if !errors.Is(err, os.ErrExist) {
			panic(fmt.Sprintf("an error '%s' occured while creating log directory", err.Error()))
		}
	}

	infoFile, err := os.Create("logs/info.log")
	if err != nil {
		panic(fmt.Sprintf("an error '%s' occured while creating info log file", err.Error()))
	}

	errorFile, err := os.Create("logs/error.log")
	if err != nil {
		panic(fmt.Sprintf("an error '%s' occured while creating error log file", err.Error()))
	}

	infoLogger := zerolog.New(infoFile).With().Timestamp().Caller().Logger()
	errorLogger := zerolog.New(errorFile).With().Timestamp().Caller().Logger()

	return &Log{
		InfoLogger:  &infoLogger,
		ErrorLogger: &errorLogger,
	}, infoFile, errorFile
}
