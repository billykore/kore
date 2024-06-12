package log

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

// Logger is service log.
type Logger struct {
	zapLogger *zap.Logger
	usecase   string
}

var (
	_once sync.Once
	_log  *Logger
)

// NewLogger return the singleton instance of Logger.
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

// Usecase set the usecase that throw the log.
func (l *Logger) Usecase(usecase string) *Logger {
	l.usecase = usecase
	return l
}

// Error log.
func (l *Logger) Error(err error) {
	l.zapLogger.Error("error", zap.String("usecase", l.usecase), zap.Error(err))
}

// Errorf is formated Error log.
func (l *Logger) Errorf(format string, a ...any) {
	l.zapLogger.Error(
		"error",
		zap.String("usecase", l.usecase),
		zap.Error(fmt.Errorf(format, a...)),
	)
}

// Fatal logs an error then shutdown the service.
func (l *Logger) Fatal(err error) {
	l.zapLogger.Fatal("fatal", zap.Error(err))
}

// Fatalf is formated Fatal log.
func (l *Logger) Fatalf(format string, a ...any) {
	l.zapLogger.Fatal(
		"fatal",
		zap.String("error", fmt.Sprintf(format, a...)),
	)
}

// Info log.
func (l *Logger) Info(v any) {
	l.zapLogger.Info(
		"info",
		zap.String("usecase", l.usecase),
		zap.Any("info", v))
}

// Infof is formated Info log.
func (l *Logger) Infof(format string, a ...any) {
	l.zapLogger.Info(
		"info",
		zap.String("usecase", l.usecase),
		zap.Any("info", fmt.Sprintf(format, a...)),
	)
}
