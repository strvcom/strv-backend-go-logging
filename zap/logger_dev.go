//go:build dev

package zap

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CreateLogger is a wrapping constructor which produces zap.Logger in a customised form based on provided config.
// Development logger provides color separated logs by log level.
func CreateLogger(config Config) (*zap.Logger, error) {
	ec := zap.NewDevelopmentEncoderConfig()
	ec.EncodeLevel = zapcore.CapitalColorLevelEncoder

	c := zap.NewDevelopmentConfig()
	c.EncoderConfig = ec
	setTimeEncoder(&c)
	prepareConfig(config, &c)

	l, err := c.Build()
	if err != nil {
		return nil, fmt.Errorf("building logger: %w", err)
	}
	return l, nil
}

// MustCreateLogger is a wrapping constructor which produces zap.Logger in a customised form based on provided config.
// Development logger provides color separated logs by log level. It panics if creation is unsuccessful.
func MustCreateLogger(config Config) *zap.Logger {
	l, err := CreateLogger(config)
	if err != nil {
		panic(err)
	}
	return l
}
