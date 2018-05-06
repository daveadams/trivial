package trivial

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type LogLevel int

const (
	DebugLevel = 1
	InfoLevel  = 0
	WarnLevel  = -1
	ErrorLevel = -2
	FatalLevel = -3
)

var levelStrings = map[LogLevel]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	FatalLevel: "FATAL",
}

type Logger struct {
	IncludeTimestamp bool
	Level            LogLevel
}

func (l *Logger) SetLogLevel(level string) {
	switch strings.ToUpper(level) {
	case "DEBUG":
		l.Level = DebugLevel
	case "INFO":
		l.Level = InfoLevel
	case "WARN":
		l.Level = WarnLevel
	case "ERROR":
		l.Level = ErrorLevel
	case "FATAL":
		l.Level = FatalLevel
	default:
		l.Level = InfoLevel
	}
}

func (l *Logger) write(level LogLevel, msg string) {
	if l.IncludeTimestamp {
		fmt.Printf("%s %s %s\n", time.Now().UTC().Format(time.RFC3339), levelStrings[level], msg)
	} else {
		fmt.Printf("%s %s\n", levelStrings[level], msg)
	}
}

func (l *Logger) Fatal(a ...interface{}) {
	l.write(FatalLevel, fmt.Sprint(a...))
	os.Exit(1)
}

func (l *Logger) Error(a ...interface{}) {
	if l.Level >= ErrorLevel {
		l.write(ErrorLevel, fmt.Sprint(a...))
	}
}

func (l *Logger) Warn(a ...interface{}) {
	if l.Level >= WarnLevel {
		l.write(WarnLevel, fmt.Sprint(a...))
	}
}

func (l *Logger) Info(a ...interface{}) {
	if l.Level >= InfoLevel {
		l.write(InfoLevel, fmt.Sprint(a...))
	}
}

func (l *Logger) Debug(a ...interface{}) {
	if l.Level >= DebugLevel {
		l.write(DebugLevel, fmt.Sprint(a...))
	}
}

func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.write(FatalLevel, fmt.Sprintf(format, a...))
	os.Exit(1)
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	if l.Level >= ErrorLevel {
		l.write(ErrorLevel, fmt.Sprintf(format, a...))
	}
}

func (l *Logger) Warnf(format string, a ...interface{}) {
	if l.Level >= WarnLevel {
		l.write(WarnLevel, fmt.Sprintf(format, a...))
	}
}

func (l *Logger) Infof(format string, a ...interface{}) {
	if l.Level >= InfoLevel {
		l.write(InfoLevel, fmt.Sprintf(format, a...))
	}
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	if l.Level >= DebugLevel {
		l.write(DebugLevel, fmt.Sprintf(format, a...))
	}
}
