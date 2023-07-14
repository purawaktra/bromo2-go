package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func AppLogLevelDetermination(value string) zerolog.Level {
	switch value {
	case "-1":
		return zerolog.TraceLevel
	case "0":
		return zerolog.DebugLevel
	case "1":
		return zerolog.InfoLevel
	case "2":
		return zerolog.WarnLevel
	case "3":
		return zerolog.ErrorLevel
	case "4":
		return zerolog.FatalLevel
	case "5":
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}

func InitLog() {
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.SetGlobalLevel(AppLogLevelDetermination(AppLogLevel))

	currentTime := time.Now()
	filename := "logs/" + AppName + "-" + currentTime.Format("02Jan2006.log")
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}

	log.Logger = log.Output(file)
}

func Log(msg string) {
	log.Log().Msg(msg)
}

// utils.Info("=== New Request ===")
func Info(msg string) {
	log.Info().Msg(msg)
}

// utils.Debug("main", "hello world")
func Debug(function string, msg interface{}) {
	log.Debug().Str("function", function).Interface("message", msg).Send()
}

// utils.DebugHTTP("GET", "/api/v1/hello-world", "hello world")
func DebugHTTP(meth string, api string, msg interface{}) {
	log.Debug().Str("method", meth).Str("api", api).Interface("body", msg).Send()
}

// utils.Warn("main", "hello world")
func Warn(function string, msg interface{}) {
	log.Warn().Str("function", function).Interface("message", msg).Send()
}

// utils.Error(errors.New("error"), "main", "hello world")
func Error(err error, function string, msg interface{}) {
	log.Error().Err(err).Str("function", function).Interface("message", msg).Send()
}

// utils.ErrorHTTP(errors.New("error"), "GET", "/api/v1/hello-world", "hello world")
func ErrorHTTP(err error, meth string, api string, msg interface{}) {
	log.Error().Err(err).Str("method", meth).Str("api", api).Interface("body", msg).Send()
}

// utils.Fatal(errors.New("error"), "main", "hello world")
func Fatal(err error, function string, msg interface{}) {
	log.Fatal().Err(err).Str("function", function).Interface("message", msg).Send()
}
