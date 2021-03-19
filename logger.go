package pastis

import "github.com/sirupsen/logrus"

func NewLogger(logFormat string) *logrus.Logger {
	logger := logrus.New()
	if logFormat == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	}
	return logger
}
