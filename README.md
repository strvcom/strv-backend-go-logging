# STRV logging
This repository contains wrappers and customisations of existing logging libraries.

## Available packages
- `zap`: A simple wrapper around [zap](https://github.com/uber-go/zap) logger that simplifies the creation of new loggers and reduces room for error.

## Examples

### Zap
```go
package main

import (
	zapx "go.strv.io/logging/zap"
)

func main() {
	logger := zapx.MustCreateLogger(zapx.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		DisableStacktrace: true,
		DisableCaller:     true,
	})
	logger.Named("main").With(
		zap.Time("current_time", time.Now().UTC()),
	).Info("app started")
}
```

Under the hood, `zap production config` is created based on the provided configuration. The resulting logger created from this config is returned.
Instead of `ts` with the current timestamp, the logger has set `ts` field with UTC time encoded in `RFC3339`. 
If an app is built with the `dev` build tag, the `CapitalColorLevelEncoder` is used. It's handy for development purposes to easily navigate in colorful log lines.
