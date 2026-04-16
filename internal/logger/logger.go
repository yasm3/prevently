package logger

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Log(level Level, msg string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] %s\n", now, level.String(), msg)
}

func (l *Logger) Debug(msg string) {
	l.Log(DEBUG, msg)
}

func (l *Logger) Info(msg string) {
	l.Log(INFO, msg)
}

func (l *Logger) Warn(msg string) {
	l.Log(WARN, msg)
}

func (l *Logger) Error(msg string) {
	l.Log(ERROR, msg)
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func (l *Logger) GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)

		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		level := INFO

		if status >= 500 {
			level = ERROR
		} else if status >= 400 {
			level = WARN
		}

		l.Log(level,
			method+" "+path+
				" | status="+itoa(status)+
				" | duration="+duration.String(),
		)
	}
}
