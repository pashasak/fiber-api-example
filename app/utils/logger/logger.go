package logger

import (
	"go.uber.org/zap"
)

var zapLog *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	zapLog = logger.Sugar()
}

type ZapWriter struct {
	Logger *zap.SugaredLogger
}

func (w ZapWriter) Write(p []byte) (n int, err error) {
	w.Logger.Info(string(p))
	return len(p), nil
}

func (w ZapWriter) Close() error {
	// modify as preferred
	return nil
}
func (w ZapWriter) Sync() error {
	w.Logger.Sync()
	return nil
}

func Info(args ...interface{}) {
	zapLog.Info(args...)
}

func Debug(args ...interface{}) {
	zapLog.Debug(args...)
}

func Error(args ...interface{}) {
	zapLog.Error(args...)
}

func Fatal(args ...interface{}) {
	zapLog.Fatal(args...)
}

func Sync() {
	zapLog.Sync()
}

func GetLogger() *zap.SugaredLogger {
	return zapLog
}
