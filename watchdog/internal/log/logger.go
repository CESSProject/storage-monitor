package log

import (
	"github.com/CESSProject/watchdog/constant"
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()
	Logger.Out = os.Stdout
	Logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: constant.TimeFormat,
		PrettyPrint:     false,
	}
}
