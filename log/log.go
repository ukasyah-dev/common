package log

import (
	"fmt"
	"os"
	"time"

	"github.com/caitlinelfring/go-env-default"
	"github.com/rs/zerolog"
	l "github.com/rs/zerolog/log"
)

func init() {
	logLevelString := env.GetDefault("LOG_LEVEL", "debug")
	logLevel, err := zerolog.ParseLevel(logLevelString)
	if err != nil {
		fmt.Printf("Invalid log level '%s', fallback to 'debug'\n", logLevelString)
		logLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(logLevel)

	l.Logger = l.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})
}

func Panic(msg string) {
	l.Panic().Msg(msg)
}

func Panicf(format string, v ...any) {
	l.Panic().Msgf(format, v...)
}

func Fatal(msg string) {
	l.Fatal().Msg(msg)
}

func Fatalf(format string, v ...any) {
	l.Fatal().Msgf(format, v...)
}

func Error(msg string) {
	l.Error().Msg(msg)
}

func Errorf(format string, v ...any) {
	l.Error().Msgf(format, v...)
}

func Warn(msg string) {
	l.Warn().Msg(msg)
}

func Warnf(format string, v ...any) {
	l.Warn().Msgf(format, v...)
}

func Info(msg string) {
	l.Info().Msg(msg)
}

func Infof(format string, v ...any) {
	l.Info().Msgf(format, v...)
}

func Debug(msg string) {
	l.Debug().Msg(msg)
}

func Debugf(format string, v ...any) {
	l.Debug().Msgf(format, v...)
}

func Trace(msg string) {
	l.Trace().Msg(msg)
}

func Tracef(format string, v ...any) {
	l.Trace().Msgf(format, v...)
}
