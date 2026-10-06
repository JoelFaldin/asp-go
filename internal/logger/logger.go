package logger

import (
	"log"
	"os"
	"time"
)

const (
	Reset  = "\033[0m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
	BgBlue = "\033[44m"
)

type Logger struct {
	Verbose   bool
	stdLogger *log.Logger
}

func New(verbose bool) *Logger {
	return &Logger{Verbose: verbose, stdLogger: log.New(os.Stdout, "", 0)}
}

func (l *Logger) Info(msg string, elapsed time.Duration) {
	l.stdLogger.Printf("%s[SUCCESS]%s %s — %s%v%s\n", Green, Reset, msg, Yellow, elapsed, Reset)
}
