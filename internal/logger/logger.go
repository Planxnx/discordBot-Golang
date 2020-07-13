package logger

import (
	"log"
	"os"
)

// NewLogger constructs a logger.
func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	return logger
}
