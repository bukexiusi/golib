package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Option func(*Logger)

// Writer
type WriterEnum int32

const (
	Console WriterEnum = iota + 1
	File
)

func Writer(writer WriterEnum) Option {
	return func(logger *Logger) {
		switch writer {
		case File:
			logger.impl.Out = rotateWriter()
		case Console:
			logger.impl.Out = os.Stdout
		default:
			panic("Invalid WriterEnum.")
		}
	}
}

// Level
type LevelEnum int32

const (
	// Error level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	Error LevelEnum = iota + 1
	// Warn level. Non-critical entries that deserve eyes.
	Warn
	// Info level. General operational entries about what's going on inside the
	// application.
	Info
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	Debug
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	Trace
)

func Level(level LevelEnum) Option {
	return func(logger *Logger) {
		switch level {
		case Error:
			logger.impl.SetLevel(logrus.ErrorLevel)
		case Warn:
			logger.impl.SetLevel(logrus.WarnLevel)
		case Info:
			logger.impl.SetLevel(logrus.InfoLevel)
		case Debug:
			logger.impl.SetLevel(logrus.DebugLevel)
		case Trace:
			logger.impl.SetLevel(logrus.TraceLevel)
		default:
			panic("Invalid LevelEnum.")
		}
	}
}
