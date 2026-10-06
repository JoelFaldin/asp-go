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

// Logger custom con colores y detalles:
func New(verbose bool) *Logger {
	return &Logger{Verbose: verbose, stdLogger: log.New(os.Stdout, "", 0)}
}

func (l *Logger) Success(msg string, elapsed time.Duration) {
	l.stdLogger.Printf("%s[SUCCESS]%s %s — %s%v%s\n", Green, Reset, msg, Yellow, elapsed, Reset)
}

func (l *Logger) Duration(msg string, totalElapsed time.Duration) {
	l.stdLogger.Printf("%s[DURATION]%s %s %v", Yellow, Reset, msg, totalElapsed)
}

func (l *Logger) Error(err error) {
	l.stdLogger.Printf("%s[ERROR]%s %v", Red, Reset, err)
}

func (l *Logger) Fatal(err error) {
	l.stdLogger.Fatalf("%s[FATAL]%s %v", Red, Reset, err)
}
