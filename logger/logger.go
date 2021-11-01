package logger

import (
	"os"
	"path"
	"time"

	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Logger struct {
	impl *logrus.Logger
}

func NewLogger(opts ...Option) *Logger {
	logger := new(Logger)

	// Default options
	logger.impl = logrus.New()
	logger.impl.Out = os.Stdout
	logger.impl.SetLevel(logrus.DebugLevel)
	logger.impl.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

	// Custom options
	for _, opt := range opts {
		opt(logger)
	}

	return logger
}

func rotateWriter() *rotate.RotateLogs {
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		cobra.CheckErr(err)
	}
	logPath := path.Join(logFilePath, "backend.log")
	writer, _ := rotate.New(
		logPath+".%Y-%m-%d %H:%M",
		rotate.WithLinkName(logPath),
		rotate.WithMaxAge(30*24*time.Hour),
		rotate.WithRotationTime(12*time.Hour),
		rotate.WithClock(rotate.Local),
	)
	return writer
}

func (l *Logger) Errorln(args ...interface{}) {
	l.impl.Errorln(args)
}

func (l *Logger) Warnln(args ...interface{}) {
	l.impl.Warnln(args)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.impl.Infoln(args)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.impl.Debugln(args)
}

func (l *Logger) Traceln(args ...interface{}) {
	l.impl.Traceln(args)
}

// Trace, Debug, Info, Warn, Error
func (l *Logger) WithMap(fields map[string]interface{}) {
	l.impl.WithFields(fields)
}
