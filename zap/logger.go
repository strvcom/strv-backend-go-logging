//go:build !dev

package zap

import (
	"fmt"

	"go.uber.org/zap"
)

// CreateLogger is a wrapping constructor which produces zap.Logger in a customised form based on provided config.
func CreateLogger(config Config) (*zap.Logger, error) {
	c := zap.NewProductionConfig()
	setTimeEncoder(&c)
	prepareConfig(config, &c)

	l, err := c.Build()
	if err != nil {
		return nil, fmt.Errorf("building logger: %w", err)
	}
	return l, nil
}

// MustCreateLogger is a wrapping constructor which produces zap.Logger in a customised form based on provided config.
// It panics if creation is unsuccessful.
func MustCreateLogger(config Config) *zap.Logger {
	l, err := CreateLogger(config)
	if err != nil {
		panic(err)
	}
	return l
}
