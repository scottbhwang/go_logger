package logger

import (
	"fmt"
	"github.com/golang/glog"
)

var depth = 2

type GlogLogger struct {
	prefix func(level LogLevel) string
}

func (e *GlogLogger) Info(a ...interface{}) {
	if ok := CheckLevel(InfoLevel); ok {
		glog.InfoDepth(depth, e.prefix(InfoLevel), fmt.Sprintln(a...))
	}
}
func (e *GlogLogger) Infof(format string, a ...interface{}) {
	if ok := CheckLevel(InfoLevel); ok {
		glog.InfoDepth(depth, fmt.Sprintf(e.prefix(InfoLevel)+format, a...))
	}
}

func (e *GlogLogger) Debug(a ...interface{}) {
	if ok := CheckLevel(DebugLevel); ok {
		glog.InfoDepth(depth, e.prefix(DebugLevel), e.prefix, fmt.Sprintln(a...))
	}
}
func (e *GlogLogger) Debugf(format string, a ...interface{}) {
	if ok := CheckLevel(DebugLevel); ok {
		glog.InfoDepth(depth, fmt.Sprintf(e.prefix(DebugLevel)+format, a...))
	}
}

func (e *GlogLogger) Error(a ...interface{}) {
	if ok := CheckLevel(ErrorLevel); ok {
		glog.ErrorDepth(depth, e.prefix(ErrorLevel), fmt.Sprintln(a...))
	}
}
func (e *GlogLogger) Errorf(format string, a ...interface{}) {
	if ok := CheckLevel(ErrorLevel); ok {
		glog.ErrorDepth(depth, fmt.Sprintf(e.prefix(ErrorLevel)+format, a...))
	}
}

func (e *GlogLogger) Fatal(a ...interface{}) {
	if ok := CheckLevel(FatalLevel); ok {
		glog.FatalDepth(depth, fmt.Sprintln(a...))
	}
}
