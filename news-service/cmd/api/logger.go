package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logsrus = logrus.New()

func init() {
	logsrus.SetFormatter(&logrus.JSONFormatter{})

	logsrus.SetLevel(logrus.InfoLevel)

	logsrus.SetOutput(os.Stdout)
}
