package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Intance logger
func Intance() *logrus.Logger {
	var logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	logger.Out, err = os.OpenFile(dir+"/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to log to file " + dir)
	}
	//defer file.Close()

	return logger
}
