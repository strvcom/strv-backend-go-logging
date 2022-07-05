//go:build dev

package zap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewLogger(t *testing.T) {
	logLevel, _ := zap.ParseAtomicLevel("warn")
	testCases := []struct {
		cfg         Config
		shouldPanic bool
	}{
		{
			cfg: Config{
				Level:       logLevel,
				Encoding:    "console",
				OutputPaths: []string{"stdout"},
			},
			shouldPanic: false,
		},
		{
			cfg:         Config{},
			shouldPanic: true,
		},
		{
			cfg: Config{
				Level:       logLevel,
				Encoding:    "slm",
				OutputPaths: []string{"stdout"},
			},
			shouldPanic: true,
		},
	}
	for _, tc := range testCases {
		if tc.shouldPanic {
			assert.Panics(t, func() {
				_ = MustCreateLogger(tc.cfg)
			})
		} else {
			l := MustCreateLogger(tc.cfg)
			l.Warn("unsaved changes", zap.String("file deleted", "srk_slm.mp4"))
		}
	}
}
