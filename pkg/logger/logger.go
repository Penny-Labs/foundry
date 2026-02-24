package logger

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerKey string

const (
	LogKey      LoggerKey = "log"
	SetupLog    LoggerKey = "setup"
	TeardownLog LoggerKey = "teardown"
)

func New(verbosity int) (logr.Logger, error) {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.DisableCaller = true
	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(-1 * verbosity))

	zapLog, err := zapConfig.Build()
	if err != nil {
		return logr.Discard(), fmt.Errorf("failed to initialize zap logger: %w", err)
	}

	return zapr.NewLogger(zapLog), nil
}

func (l LoggerKey) String() string {
	return string(l)
}

// LogFromContextOrDiscard will check the LogKey if the logger exists in the context.
// WHen the logger does not exist in the context we will create a discard logger from zap and it will not output anything
func LogFromContextOrDiscard(ctx context.Context, logKey LoggerKey) logr.Logger {
	if v, ok := ctx.Value(logKey).(logr.Logger); ok {
		return v
	}

	return zapr.NewLogger(zap.NewNop())
}
