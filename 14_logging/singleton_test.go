package logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	// without logrus.New()
	// this error it will expose to global global
	logrus.Info("Hello world")
}
