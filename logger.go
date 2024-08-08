package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Version represents the current version of the application.
const Version = "1.0.1"

// Loglevel represents the severity of the log message.
type Loglevel string

const (
	// INFO represents informational messages that highlight the progress of the application at coarse-grained level.
	INFO Loglevel = "INFO"
	// WARN represents potentially harmful situations.
	WARN Loglevel = "WARN"
	// ERROR represents error events that might still allow the application to continue running.
	ERROR Loglevel = "ERROR"
	// DEBUG represents fine-grained informational events that are most useful to debug an application.
	DEBUG Loglevel = "DEBUG"
)

// LogLevelOrder maps each log level to an integer for comparison purposes.
var LogLevelOrder = map[Loglevel]int{
	INFO:  0,
	WARN:  1,
	ERROR: 2,
	DEBUG: 3,
}

// Logger is a custom logger that includes a log level.
type Logger struct {
	*log.Logger
	loglevel Loglevel
}

// loggerTimeStamp returns the current time formatted as a string.
func LoggerTimeStamp() string {
	return time.Now().Format("2006/01/02-15:04:05")
}

// NewLogger creates a new Logger instance with the default log level set to INFO.
func NewLogger() *Logger {
	newLogger := &Logger{
		log.New(os.Stdout, "", log.LstdFlags),
		INFO,
	}
	newLogger.Info(fmt.Sprintf("Logger initialized with log level %s", newLogger.loglevel))
	newLogger.Debug("!!! Debug log level enabled !!!")
	return newLogger
}

// ShouldBeLogged determines if a message with the given log level should be logged.
func (l *Logger) ShouldBeLogged(loglevel Loglevel) bool {
	return l.ValidateLogLevel(loglevel) && LogLevelOrder[l.loglevel] >= LogLevelOrder[loglevel]
}

// ValidateLogLevel checks if the provided log level is valid.
func (l *Logger) ValidateLogLevel(loglevel Loglevel) bool {
	_, ok := LogLevelOrder[loglevel]
	return ok
}

// Write logs a message with the given log level if it meets the current log level threshold.
func (l *Logger) Write(loglevel Loglevel, msg string, v ...any) {
	if l.ShouldBeLogged(loglevel) {
		l.Writer().Write([]byte(
			fmt.Sprintf("[LOG-%s][%v] %s\n", loglevel, LoggerTimeStamp(), fmt.Sprintf(msg, v...)),
		))
	}
}

// Info logs a message at the INFO log level.
func (l *Logger) Info(msg string, v ...any) {
	l.Write(INFO, msg, v...)
}

// Warn logs a message at the WARN log level.
func (l *Logger) Warn(msg string, v ...any) {
	l.Write(WARN, msg, v...)
}

// Error logs a message at the ERROR log level.
func (l *Logger) Error(msg string, v ...any) {
	l.Write(ERROR, msg, v...)
}

// Debug logs a message at the DEBUG log level.
func (l *Logger) Debug(msg string, v ...any) {
	l.Write(DEBUG, msg, v...)
}

// SetLogLevel sets the log level for the logger.
func (l *Logger) SetLogLevel(loglevel Loglevel) {
	if l.ValidateLogLevel(loglevel) {
		l.loglevel = loglevel
		l.Info(fmt.Sprintf("Log level set to %s", loglevel))
		return
	}
	l.Error(fmt.Sprintf("Invalid log level %s! Loglevel still is set to %s", loglevel, l.loglevel))
}

// GetLogLevel returns the current log level of the logger.
func (l *Logger) GetLogLevel() Loglevel {
	return l.loglevel
}

// Default is the default Logger instance with the default log level set to INFO.
var Default = NewLogger()
