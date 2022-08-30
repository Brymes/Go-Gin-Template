package config

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strings"
)

func CreateLogger(fileName string) *logrus.Logger {
	//Initialize Logger Instance to file or external Service
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	return logger
}

func InitRequestLogger(service string) (*bytes.Buffer, *log.Logger) {
	// Initialize Logger Instance for Each Request so as to have clean logs
	b := &bytes.Buffer{}
	prefix := fmt.Sprintf("%v: ", strings.ToUpper(service))
	reqlogger := log.New(b, prefix, log.Ltime|log.Lshortfile)

	return b, reqlogger
}
