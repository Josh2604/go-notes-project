package zlog

import (
	"bytes"
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

func New(isDebug bool) *Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	return &Logger{&logger}
}

func NewConsole(isDebug bool) *Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &Logger{&logger}
}

type Tags map[string]interface{}

func (l *Logger) Info(msg string, tags ...string) {
	l.Logger.Info().Msgf(msg, addFields(tags...))
}

func (l *Logger) Error(msg string, err error, tags Tags) {
	l.Logger.Error().Err(err).Str("tags", parseKeyValues(tags)).Msg(msg)
}

func addFields(values ...string) string {
	fields := ""
	for _, value := range values {
		fields += " " + value + " - "
	}
	return fields
}

func parseKeyValues(m map[string]interface{}) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s:%s,", key, value)
	}
	return b.String()
}
