package logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	// logger.SetFormatter(&logrus.JSONFormatter{})
	logger.Trace("Info")
	logger.Debug("Info")
	logger.Info("Info")
	logger.Warn("Info")
	logger.Error("Info")
}
