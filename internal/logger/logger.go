package logger

import (
	"os"
	"github.com/rs/zerolog"
)

func Logger() zerolog.Logger{
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	return logger
}