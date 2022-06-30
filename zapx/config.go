package zapx

import (
	"go.uber.org/zap"
)

type Config struct {
	// Level is the minimum enabled logging level. Note that this is a dynamic
	// level, so calling Config.Level.SetLevel will atomically change the log
	// level of all loggers descended from this config.
	Level zap.AtomicLevel `json:"level"`
	// DisableCaller stops annotating logs with the calling function's file
	// name and line number. By default, all logs are annotated.
	DisableCaller bool `json:"disable_caller"`
	// DisableStacktrace completely disables automatic stacktrace capturing. By
	// default, stacktraces are captured for WarnLevel and above logs in
	// development and ErrorLevel and above in production.
	DisableStacktrace bool `json:"disable_stacktrace"`
	// Encoding sets the logger's encoding. Valid values are "json" and
	// "console", as well as any third-party encodings registered via
	// RegisterEncoder.
	Encoding string `json:"encoding"`
	// OutputPaths is a list of URLs or file paths to write logging output to.
	// See Open for details.
	OutputPaths []string `json:"output_paths"`
	// ErrorOutputPaths is a list of URLs to write internal logger errors to.
	// The default is standard error.
	//
	// Note that this setting only affects internal errors; for sample code that
	// sends error-level logs to a different location from info- and debug-level
	// logs, see the package-level AdvancedConfiguration example.
	ErrorOutputPaths []string `json:"error_output_paths"`
	// InitialFields is a collection of fields to add to the root logger.
	InitialFields map[string]interface{} `json:"initial_fields"`
}

func prepareConfig(src Config, dest *zap.Config) error {
	dest.Level = src.Level

	dest.DisableCaller = src.DisableCaller
	dest.DisableStacktrace = src.DisableStacktrace

	if src.Encoding != "" {
		dest.Encoding = src.Encoding
	}

	if len(src.OutputPaths) != 0 {
		dest.OutputPaths = src.OutputPaths
	}

	if len(src.ErrorOutputPaths) != 0 {
		dest.ErrorOutputPaths = src.ErrorOutputPaths
	}

	if len(src.InitialFields) != 0 {
		dest.InitialFields = src.InitialFields

	}

	return nil
}
