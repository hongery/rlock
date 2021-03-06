package rlock

import (
	"github.com/hunterhug/marmot/util/go-logging"
	"os"
)

// Global logger config for debug
var (
	Logger = logging.MustGetLogger("RLock")

	format = logging.MustStringFormatter(
		"%{color}%{time:2006-01-02 15:04:05.000} %{longpkg}:%{longfunc} [%{level:.5s}]:%{color:reset} %{message}",
	)

	// Level name you can refer
	LevelNames = []string{
		"CRITICAL",
		"ERROR",
		"WARNING",
		"NOTICE",
		"INFO",
		"DEBUG",
	}
)

// Init log record
func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
	logging.SetLevel(logging.INFO, "RLock")
}

// Set log level
func SetLogLevel(level string) {
	lvl, _ := logging.LogLevel(level)
	logging.SetLevel(lvl, "RLock")
}

// Return global log
func Log() *logging.Logger {
	return Logger
}
