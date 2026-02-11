package logger

import (
	"fmt"
	"os"
	"time"
)

// Logger provides structured logging
type Logger struct {
	verbose bool
}

// New creates a new logger
func New(verbose bool) *Logger {
	return &Logger{
		verbose: verbose,
	}
}

// Info logs an info message
func (l *Logger) Info(format string, args ...interface{}) {
	l.log("INFO", format, args...)
}

// Debug logs a debug message (only if verbose)
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.verbose {
		l.log("DEBUG", format, args...)
	}
}

// Success logs a success message
func (l *Logger) Success(format string, args ...interface{}) {
	l.log("SUCCESS", format, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(format string, args ...interface{}) {
	l.log("WARN", format, args...)
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	l.log("ERROR", format, args...)
}

// log formats and outputs a log message
func (l *Logger) log(level, format string, args ...interface{}) {
	timestamp := time.Now().Format("15:04:05")
	message := fmt.Sprintf(format, args...)

	var prefix string
	switch level {
	case "INFO":
		prefix = "ℹ️"
	case "DEBUG":
		prefix = "🔍"
	case "SUCCESS":
		prefix = "✅"
	case "WARN":
		prefix = "⚠️"
	case "ERROR":
		prefix = "❌"
	default:
		prefix = "✨"
	}

	output := fmt.Sprintf("[%s] %s %s\n", timestamp, prefix, message)

	if level == "ERROR" {
		fmt.Fprint(os.Stderr, output)
	} else {
		fmt.Fprint(os.Stdout, output)
	}
}

// SetVerbose sets the verbose flag
func (l *Logger) SetVerbose(verbose bool) {
	l.verbose = verbose
}

// IsVerbose returns the verbose flag
func (l *Logger) IsVerbose() bool {
	return l.verbose
}
