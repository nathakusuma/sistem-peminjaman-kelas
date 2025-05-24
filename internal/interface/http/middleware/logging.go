package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/ctxkey"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/log"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Generate trace ID
		traceID := uuid.New()
		ctx := context.WithValue(r.Context(), ctxkey.TraceID, traceID)
		r = r.WithContext(ctx)

		// Create response writer wrapper to capture status code
		wrapper := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}

		// Process request
		next.ServeHTTP(wrapper, r)

		// Log response
		duration := time.Since(start)
		log.Info(ctx).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Str("remote_addr", r.RemoteAddr).
			Str("user_agent", r.UserAgent()).
			Int("status_code", wrapper.statusCode).
			Dur("duration", duration).
			Msg("Request completed")
	})
}

// Response writer wrapper to capture status code
type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}
