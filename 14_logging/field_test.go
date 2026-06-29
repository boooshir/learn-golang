package logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {

	logger := logrus.New()

	logger.WithField("username", "bangsyir").Info("Hello worlds")

	logger.WithFields(logrus.Fields{
		"username": "boooshir",
		"name":     "eko Kurniawan",
	}).Infof("Hello worlds")
}
