package logger

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type Logger struct {
	logger logrus.FieldLogger
}

var logrusLogger = logrus.New()
var Log = NewLogger(logrusLogger)

func NewLogger(logger logrus.FieldLogger) *Logger {
	return &Logger{logger}
}

func (l *Logger) Info(data string) {
	l.logger.Infof(time.Now().UTC().Format(time.RFC3339) + ":" + data + "\n")
}

func (l *Logger) Fatal(data string) {
	l.logger.Fatalf(time.Now().UTC().Format(time.RFC3339) + ":" + data + "\n")
}

func (l *Logger) Error(data string) {
	l.logger.Error(time.Now().UTC().Format(time.RFC3339) + ":" + data + "\n")
}

func (l *Logger) ErrorHttp(req *http.Request, data string, code int) {
	l.logger.Error(req.Method + " " + req.RequestURI + " " + strconv.Itoa(code) + ": " + data + "\n")
}
