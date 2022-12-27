package log

import (
	"go.uber.org/zap"
)

type Interface interface {
	Fatal(obj interface{})
	Debug(obj interface{})
	Error(obj interface{})
	Info(obj interface{})
}

type logger struct {
	log *zap.SugaredLogger
}

func Init() Interface {
	log, _ := zap.NewDevelopment()
	defer log.Sync()

	sugar := log.Sugar()

	return &logger{log: sugar}
}

func (l *logger) Fatal(obj interface{}) {
	l.log.Fatal(obj)
}

func (l *logger) Debug(obj interface{}) {
	l.log.Debug(obj)
}

func (l *logger) Error(obj interface{}) {
	l.log.Error(obj)
}

func (l *logger) Info(obj interface{}) {
	l.log.Info(obj)
}
