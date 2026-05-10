package logger

import (
	"os"
	"time"
	"path/filepath"
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

func Logger() zerolog.Logger{
	logger := zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		NoColor: true,
		FormatLevel: func(i interface{}) string {return strings.ToUpper(fmt.Sprintf("[%s]", i))},
		FormatMessage: func(i interface{}) string {return fmt.Sprintf("| %s |", i)},
		FormatCaller: func(i interface{}) string {return filepath.Base(fmt.Sprintf("%s", i))},
		PartsExclude: []string{zerolog.TimestampFieldName}}).With().Timestamp().Caller().Logger()
	return logger
}