package clog

import "go.uber.org/zap"

type SugarLogger struct {
	zapSugar *zap.SugaredLogger
}

func (l *SugarLogger) Debug(msg string, keyValuePairs ...interface{}) {
	l.zapSugar.Debugw(msg, keyValuePairs...)
}

func (l *SugarLogger) Info(msg string, keyValuePairs ...interface{}) {
	l.zapSugar.Infow(msg, keyValuePairs...)
}

func (l *SugarLogger) Error(msg string, keyValuePairs ...interface{}) {
	l.zapSugar.Errorw(msg, keyValuePairs...)
}

func (l *SugarLogger) Warn(msg string, keyValuePairs ...interface{}) {
	l.zapSugar.Warnw(msg, keyValuePairs...)
}