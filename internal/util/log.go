// Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/union-jd-go.

package util

import (
	"context"
	"time"

	"go.uber.org/zap"
)

// ILogger is a logger interface.
type ILogger interface {
	// Debugf logs a message at debug level.
	Debugf(ctx context.Context, format string, args ...interface{})
	// Infof logs a message at info level.
	Infof(ctx context.Context, format string, args ...interface{})
	// Warnf logs a message at warn level.
	Warnf(ctx context.Context, format string, args ...interface{})
	// Errorf logs a message at error level.
	Errorf(ctx context.Context, format string, args ...interface{})
	// Fatalf logs a message at fatal level.
	Fatalf(ctx context.Context, format string, args ...interface{})
}

// DefaultLogger is a default logger.
type DefaultLogger struct {
	ctx    context.Context
	logger ILogger
}

// Debugf logs a message at debug level.
func (d *DefaultLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	// TODO
}

// Infof logs a message at info level.
func (d *DefaultLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	// TODO
}

// Warnf logs a message at warn level.
func (d *DefaultLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	// TODO
}

// Errorf logs a message at error level.
func (d *DefaultLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	// TODO
}

// Fatalf logs a message at fatal level.
func (d *DefaultLogger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	// TODO
}

// NewLogger returns a new Logger.
func NewLogger() ILogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", ""),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	return &DefaultLogger{}
}
