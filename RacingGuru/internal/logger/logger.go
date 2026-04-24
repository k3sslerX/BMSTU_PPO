package logger

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"strings"
)

type Logger struct {
	base *slog.Logger
}

func New(level string, output io.Writer) (*Logger, error) {
	parsedLevel, err := parseLevel(level)
	if err != nil {
		return nil, err
	}

	return &Logger{
		base: slog.New(slog.NewTextHandler(output, &slog.HandlerOptions{
			Level: parsedLevel,
		})),
	}, nil
}

func (l *Logger) Debug(msg string, args ...any) {
	l.base.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.base.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.base.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.base.Error(msg, args...)
}

func (l *Logger) Debugf(format string, args ...any) {
	l.base.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Infof(format string, args ...any) {
	l.base.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Warnf(format string, args ...any) {
	l.base.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Errorf(format string, args ...any) {
	l.base.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) StdLogger() *log.Logger {
	return slog.NewLogLogger(l.base.Handler(), slog.LevelError)
}

func parseLevel(level string) (slog.Level, error) {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		return slog.LevelDebug, nil
	case "", "info":
		return slog.LevelInfo, nil
	case "warn", "warning":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return 0, fmt.Errorf("unsupported log level: %q", level)
	}
}
