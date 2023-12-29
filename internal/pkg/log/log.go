package log

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

type Logger struct {
	zapLogger *zap.Logger
	usecase   string
}

var (
	_once sync.Once
	_log  *Logger
)

func NewLogger() *Logger {
	_once.Do(func() {
		zapLogger := newLogger()
		_log = &Logger{zapLogger: zapLogger}
	})
	return _log
}

func newLogger() *zap.Logger {
	zapLogger, _ := zap.NewProduction()
	_ = zapLogger.Sync()
	return zapLogger
}

func (l *Logger) Usecase(usecase string) *Logger {
	l.usecase = usecase
	return l
}

func (l *Logger) Error(err error) {
	l.zapLogger.Error("error", zap.String("usecase", l.usecase), zap.Error(err))
}

func (l *Logger) Fatal(err error) {
	l.zapLogger.Fatal("fatal", zap.Error(err))
}

func (l *Logger) Fatalf(format string, a ...any) {
	l.zapLogger.Fatal(
		"fatal",
		zap.String("error", fmt.Sprintf(format, a...)),
	)
}

func (l *Logger) Infof(format string, a ...any) {
	l.zapLogger.Info(
		"info",
		zap.String("usecase", l.usecase),
		zap.Any("info", fmt.Sprintf(format, a...)),
	)
}
