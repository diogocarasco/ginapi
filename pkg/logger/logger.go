package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func Init() {
	Logger = logrus.New()
	Logger.Out = os.Stdout
	Logger.SetLevel(logrus.InfoLevel)
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}
