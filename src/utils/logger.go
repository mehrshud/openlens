package utils

import (
	"log"
)

// LoggerConfig represents the configuration for the logger utility
type LoggerConfig struct {
	Debug bool
}

// Logger represents the logger utility
type Logger struct {
	config LoggerConfig
}

// NewLogger returns a new instance of the logger utility
func NewLogger(config LoggerConfig) *Logger {
	return &Logger{config: config}
}

// Debug logs a debug message if the debug flag is enabled
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.config.Debug {
		log.Printf("[DEBUG] "+format, v...)
	}
}

// Info logs an info message
func (l *Logger) Info(format string, v ...interface{}) {
	log.Printf("[INFO] "+format, v...)
}

// Warn logs a warning message
func (l *Logger) Warn(format string, v ...interface{}) {
	log.Printf("[WARN] "+format, v...)
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	log.Printf("[ERROR] "+format, v...)
}

// Fatal logs a fatal message and exits the program
func (l *Logger) Fatal(format string, v ...interface{}) {
	log.Printf("[FATAL] "+format, v...)
}

// LogUser logs user information
func (l *Logger) LogUser(user User) {
	l.Info("User ID: %d, Name: %s, Email: %s", user.ID, user.Name, user.Email)
}

// LogDashboard logs dashboard information
func (l *Logger) LogDashboard(dashboard Dashboard) {
	l.Info("Dashboard ID: %d, Name: %s", dashboard.ID, dashboard.Name)
	for _, visualization := range dashboard.Visualizations {
		l.Info("Visualization ID: %d, Name: %s, Type: %s", visualization.ID, visualization.Name, visualization.Type)
	}
}
