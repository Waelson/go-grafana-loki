package log

import (
	"context"
	"github.com/Waelson/go-grafana-loki/internal/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger() *ZapLogger {
	// Configuração personalizada para o encoder
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalLevelEncoder,                                 // Loga o nível de log em maiúsculas
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.0000000Z"), // Formato de timestamp desejado
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Criação do core com o encoder JSON e nível de log INFO
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),    // Usa encoder JSON com config customizada
		zapcore.AddSync(zapcore.Lock(os.Stdout)), // Envia os logs para o stdout
		zapcore.InfoLevel,                        // Nível de log mínimo
	)

	// Retorna um novo logger com o core personalizado
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &ZapLogger{logger: zapLogger}
}

// Métodos sem contexto
func (z *ZapLogger) Info(message string, fields ...Field) {
	z.logger.Info(message, convertFields(fields)...)
}

func (z *ZapLogger) Warn(message string, fields ...Field) {
	z.logger.Warn(message, convertFields(fields)...)
}

func (z *ZapLogger) Error(message string, fields ...Field) {
	z.logger.Error(message, convertFields(fields)...)
}

func (z *ZapLogger) Debug(message string, fields ...Field) {
	z.logger.Debug(message, convertFields(fields)...)
}

// Métodos com contexto
func (z *ZapLogger) InfoCtx(ctx context.Context, message string, fields ...Field) {
	z.logger.Info(message, append(convertFields(fields), extractContextFields(ctx)...)...)
}

func (z *ZapLogger) WarnCtx(ctx context.Context, message string, fields ...Field) {
	z.logger.Warn(message, append(convertFields(fields), extractContextFields(ctx)...)...)
}

func (z *ZapLogger) ErrorCtx(ctx context.Context, message string, fields ...Field) {
	z.logger.Error(message, append(convertFields(fields), extractContextFields(ctx)...)...)
}

func (z *ZapLogger) DebugCtx(ctx context.Context, message string, fields ...Field) {
	z.logger.Debug(message, append(convertFields(fields), extractContextFields(ctx)...)...)
}

// With permite criar um logger com campos adicionais de contexto.
func (z *ZapLogger) With(fields ...Field) Logger {
	return &ZapLogger{logger: z.logger.With(convertFields(fields)...)}
}

// convertFields converte os campos da abstração Field para o formato zap.Field.
func convertFields(fields []Field) []zap.Field {
	var zapFields []zap.Field
	for _, field := range fields {
		zapFields = append(zapFields, zap.Any(field.Key, field.Value))
	}
	return zapFields
}

// extractContextFields extrai informações do contexto e as converte em zap.Field.
func extractContextFields(ctx context.Context) []zap.Field {
	var fields []zap.Field
	if traceID, ok := ctx.Value(middleware.TraceIDKey).(string); ok {
		fields = append(fields, zap.String("trace_id", traceID))
	}
	if requestID, ok := ctx.Value(middleware.RequestIDKey).(string); ok {
		fields = append(fields, zap.String("request_id", requestID))
	}

	if requestID, ok := ctx.Value(middleware.InstanceIDKey).(string); ok {
		fields = append(fields, zap.String("instance_id", requestID))
	}

	if requestID, ok := ctx.Value(middleware.SpanIDKey).(string); ok {
		fields = append(fields, zap.String("span_id", requestID))
	}

	return fields
}
