package types

import "go.uber.org/zap"

type Logger interface {
	DPanic(...interface{})
	DPanicf(string, ...interface{})
	DPanicw(string, ...interface{})
	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugw(string, ...interface{})
	Desugar() *zap.Logger
	Error(...interface{})
	Errorf(string, ...interface{})
	Errorw(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalw(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Infow(string, ...interface{})
	Named(string) *zap.SugaredLogger
	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicw(string, ...interface{})
	Sync() error
	Warn(...interface{})
	Warnf(string, ...interface{})
	Warnw(string, ...interface{})
	With(...interface{}) *zap.SugaredLogger
}
