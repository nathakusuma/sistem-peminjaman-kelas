package log

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/ctxkey"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/enum"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/infrastructure/config"
)

var (
	logger zerolog.Logger
	once   sync.Once
)

func NewLogger() *zerolog.Logger {
	once.Do(func() {
		zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
			fn := runtime.FuncForPC(pc)
			if fn == nil {
				return fmt.Sprintf("%s:%d", filepath.Base(file), line)
			}
			return fmt.Sprintf("%s:%d %s", filepath.Base(file), line, filepath.Base(fn.Name()))
		}

		writers := []io.Writer{
			zerolog.ConsoleWriter{
				Out:        os.Stderr,
				TimeFormat: time.RFC3339,
			},
		}

		// Add file output only if not in test environment
		if config.GetEnv().Env != enum.EnvTesting {
			fileWriter := &lumberjack.Logger{
				Filename:   fmt.Sprintf("./storage/logs/app-%s.log", time.Now().Format("2006-01-02")),
				LocalTime:  true,
				Compress:   true,
				MaxSize:    100, // megabytes
				MaxAge:     7,   // days
				MaxBackups: 3,
			}
			writers = append(writers, fileWriter)
		}

		logger = zerolog.New(zerolog.MultiLevelWriter(writers...)).
			With().
			Timestamp().
			Logger()
	})

	return &logger
}

// addContextFields extracts context values and adds them to the event
func addContextFields(ctx context.Context, event *zerolog.Event) *zerolog.Event {
	if ctx == nil {
		return event
	}

	// Extract user ID from context if it exists
	if userID := ctx.Value(ctxkey.UserEmail); userID != nil {
		event = event.Interface("requester.id", userID)
	}

	// Extract trace ID from context if it exists
	if traceID := ctx.Value(ctxkey.TraceID); traceID != nil {
		event = event.Interface("trace_id", traceID)
	}

	return event
}

func Debug(ctx context.Context) *zerolog.Event {
	event := logger.Debug()
	if !event.Enabled() {
		return event
	}

	event = event.Caller(1)
	return addContextFields(ctx, event)
}

// Usage: log.Info(ctx).Str("key", "value").Msg("info message")
func Info(ctx context.Context) *zerolog.Event {
	event := logger.Info()
	if !event.Enabled() {
		return event
	}

	event = event.Caller(1)
	return addContextFields(ctx, event)
}

func Warn(ctx context.Context) *zerolog.Event {
	event := logger.Warn()
	if !event.Enabled() {
		return event
	}

	event = event.Caller(1)
	return addContextFields(ctx, event)
}

func Error(ctx context.Context) *zerolog.Event {
	event := logger.Error()
	if !event.Enabled() {
		return event
	}

	event = event.Caller(1)
	return addContextFields(ctx, event)
}

func Fatal(ctx context.Context) *zerolog.Event {
	event := logger.Fatal().Caller(1)
	return addContextFields(ctx, event)
}

func Panic(ctx context.Context) *zerolog.Event {
	event := logger.Panic().Caller(1)
	return addContextFields(ctx, event)
}
