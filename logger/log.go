package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)


var Log *logrus.Logger


func InitLogger(logLevel string) {
	level, _ := logrus.ParseLevel(logLevel)
	Log = &logrus.Logger{
		Out:       os.Stdout,
		Level:     level,
		Formatter: &logrus.JSONFormatter{},
	}
}
