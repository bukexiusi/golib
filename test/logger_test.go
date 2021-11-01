package test

import (
	"testing"

	"github.com/bukexiusi/golib/logger"
)

// TestLogger
func TestLogger(t *testing.T) {
	log := logger.NewLogger(logger.Writer(logger.File), logger.Level(logger.Info))
	log.Infoln("1", 2)
}
