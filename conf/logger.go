package conf

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	TIME_FORMAT = "2006/01/02 15:04:05"
)

func InitLogger() {
	switch Conf.Log.LogLevel {
	case "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	case "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "INFO":
		logrus.SetLevel(logrus.InfoLevel)
	case "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	case "PANIC":
		logrus.SetLevel(logrus.PanicLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	switch Conf.App.Mode {
	case "Development":
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
			TimestampFormat: TIME_FORMAT,
		})
		logrus.Warnln("Run in Development, don't use this in product environment !!!")
		logrus.SetOutput(os.Stdout)
	case "Production":
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: TIME_FORMAT,
		})
		logrus.Infoln("Run in Production mode.")
		logrus.Infof("Log file save in %s\n", Conf.Log.logFile)
		logger := &lumberjack.Logger{
			Filename:   Conf.Log.logFile,
			MaxSize:    Conf.Log.logFileMaxSize,
			MaxBackups: Conf.Log.logFileMaxNums,
			MaxAge:     Conf.Log.logFileMaxKeepTime,
			Compress:   Conf.Log.logFileCompress,
		}
		logrus.SetOutput(logger)
	}
}
