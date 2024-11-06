package middleware

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"os"
)

// Context keys para os identificadores
type contextKey string

const (
	TraceIDKey    contextKey = "trace_id"
	RequestIDKey  contextKey = "request_id"
	InstanceIDKey contextKey = "instance_id"
	SpanIDKey     contextKey = "span_id"
)

// NewContextMiddleware cria um middleware que injeta trace_id, request_id, instance_id e span_id no contexto
func NewContextMiddleware() func(next http.Handler) http.Handler {
	// Obtém o hostname da máquina como instance_id
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown-instance" // Valor padrão caso não seja possível obter o hostname
	} else {
		hostname = fmt.Sprintf("i-%s", hostname)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Gera novos valores para trace_id e span_id
			traceID := uuid.New().String()
			spanID := uuid.New().String()

			// Obtém o request_id da requisição, se disponível (pode vir de um header customizado)
			requestID := r.Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = uuid.New().String()
			}

			// Cria um novo contexto com os valores injetados
			ctx := context.WithValue(r.Context(), TraceIDKey, traceID)
			ctx = context.WithValue(ctx, RequestIDKey, requestID)
			ctx = context.WithValue(ctx, InstanceIDKey, hostname)
			ctx = context.WithValue(ctx, SpanIDKey, spanID)

			// Chama o próximo handler com o novo contexto
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
