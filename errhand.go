package errhand

import (
	"github.com/kattaris/errhand/v2/internal/logger"
)

type Handler struct {
}

// HandleSimpleErr handles simple errors
func (Handler) HandleSimpleErr(err error, message string) {
	if err != nil {
		logger.Error(message, err)
	}
}

// Return Handler with log path and level
func New(sysVarLogPath string, level string) Handler {
	logger.SetPath(sysVarLogPath)
	logger.SetLevel(level)
	return Handler{}
}

// Print with new line
func (Handler) Println(v ...interface{}) {
	logger.Println(v...)
}

// Print with format
func (Handler) Printf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}
