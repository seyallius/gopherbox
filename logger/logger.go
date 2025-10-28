// Package logger. logger provides high-performance structured logging using Uber's Zap.
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ------------------------------------- Types -------------------------------------

type logLevel string

type logPrintingTyp string

// ------------------------------------- Constants & Variables -------------------------------------

const (
	logLevelDev  logLevel = "development"
	logLevelProd logLevel = "production"

	printTypeJson      logPrintingTyp = "json"
	printTypeFormatted logPrintingTyp = "formated"
)

// ------------------------------------- Public functions -------------------------------------

// NewLogger creates a new logger instance with the specified log level and log printing type.
// It returns a pointer to the logger and a pointer to the sugared logger.
func NewLogger(level logLevel, logPrintingType string) (*zap.Logger, *zap.SugaredLogger) {
	log := initializeLogger(level, logPrintingType)
	sugar := log.Sugar()
	return log, sugar
}

// ------------------------------------- Private helper functions -------------------------------------

// initializeLogger creates and configures the Zap logger instance
func initializeLogger(level logLevel, printType string) *zap.Logger {
	var logger *zap.Logger
	var err error

	if printType == "" {
		printType = string(printTypeFormatted)
	}

	// Choose base config
	var cfg zap.Config
	switch level {
	case logLevelProd:
		cfg = zap.NewProductionConfig()
	default:
		cfg = zap.NewDevelopmentConfig()
	}

	// Override encoding
	switch logPrintingTyp(printType) {
	case printTypeJson:
		cfg.Encoding = "json"
	default:
		cfg.Encoding = "console"
	}

	// Parse log level
	if level != "" {
		var lvl zapcore.Level
		if err := lvl.Set(string(level)); err == nil {
			cfg.Level = zap.NewAtomicLevelAt(lvl)
		}
	}

	logger, err = cfg.Build()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}

	return logger
}
