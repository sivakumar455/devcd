package logger

import (
	"log/slog"
	"os"
)

// Logger interface
type Logger interface {
	Info(msg string, args ...any)
	Debug(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

var log Logger

func init() {
	// Set logging level, format, and destination from environment variables or default values
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	logFormat := os.Getenv("LOG_FORMAT")
	if logFormat == "" {
		logFormat = "text"
	}
	logDestination := os.Getenv("LOG_DESTINATION")
	if logDestination == "" {
		logDestination = "stdout"
	}

	InitLogger(logLevel, logFormat, logDestination)

	//fmt.Println("LOG_LEVEL: ", logLevel)
	Debug("Default Logging level set", "logLevel", logLevel)

}

func InitLogger(level, format, destination string) {

	var logLevel = getLogLevel(level)

	var handler slog.Handler
	switch format {
	case "json":
		handler = slog.NewJSONHandler(getLogDestination(destination), &slog.HandlerOptions{
			Level: logLevel,
		})
	case "text":
		handler = slog.NewTextHandler(getLogDestination(destination), &slog.HandlerOptions{
			Level: logLevel,
		})
	default:
		handler = slog.NewTextHandler(getLogDestination(destination), &slog.HandlerOptions{
			Level: logLevel,
		})
	}
	log = slog.New(handler)
}

func getLogLevel(level string) slog.Level {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	return logLevel
}

func getLogDestination(destination string) *os.File {
	switch destination {
	case "stdout":
		return os.Stdout
	case "stderr":
		return os.Stderr
	default:
		file, err := os.OpenFile(destination, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		return file
	}
}

// SetLogger allows setting a custom logger (e.g., a mock logger)
func SetLogger(customLogger Logger) {
	log = customLogger
}

func SetLogLevel(level string) {
	Info("Changing Log level", "logLevel", level)
	InitLogger(level, "text", "stdout")
}

// Exported functions for logging
func Info(msg string, args ...any) {
	log.Info(msg, args...)
}

func Debug(msg string, args ...any) {
	log.Debug(msg, args...)
}

func Warn(msg string, args ...any) {
	log.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	log.Error(msg, args...)
}
