package logger

import (
	"fmt"
	"strings"
)

type LogLevel int

const (
	TraceLevel LogLevel = iota
	DebugLevel
	InfoLevel
	ErrorLevel
	FatalLevel
)

func (e LogLevel) Set(s string) error {
	fmt.Println("日志等级设为:", ParseLevel(s))
	return nil
}

func (e LogLevel) String() string {
	switch e {
	case TraceLevel:
		return "[TRACE]"
	case DebugLevel:
		return "[DEBUG]"
	case InfoLevel:
		return "[INFO]"
	case ErrorLevel:
		return "[ERROR]"
	case FatalLevel:
		return "[FATAL]"
	default:
		return "[NIL]"
	}
}

func ParseLevel(str string) LogLevel {
	switch {
	case strings.EqualFold(str, TraceLevel.String()):
		return TraceLevel
	case strings.EqualFold(str, DebugLevel.String()):
		return DebugLevel
	case strings.EqualFold(str, InfoLevel.String()):
		return InfoLevel
	case strings.EqualFold(str, ErrorLevel.String()):
		return ErrorLevel
	case strings.EqualFold(str, FatalLevel.String()):
		return FatalLevel
	default:
		return InfoLevel
	}
}
