package log

import "context"

// Logger é uma interface para abstrair o framework de log.
type Logger interface {
	Info(message string, fields ...Field)
	Warn(message string, fields ...Field)
	Error(message string, fields ...Field)
	Debug(message string, fields ...Field)

	InfoCtx(ctx context.Context, message string, fields ...Field)
	WarnCtx(ctx context.Context, message string, fields ...Field)
	ErrorCtx(ctx context.Context, message string, fields ...Field)
	DebugCtx(ctx context.Context, message string, fields ...Field)

	With(fields ...Field) Logger
}

type Field struct {
	Key   string
	Value interface{}
}
