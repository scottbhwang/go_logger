package logger

import (
	"fmt"
	"github.com/golang/glog"
)

var depth = 2

type GlogLogger struct {
	prefix string
}

func (e *GlogLogger) Info(a ...interface{}) {
	if ok := CheckLevel(InfoLevel); ok {
		glog.InfoDepth(depth, e.prefix, fmt.Sprint(a...))
	}
}
func (e *GlogLogger) Infof(format string, a ...interface{}) {
	if ok := CheckLevel(InfoLevel); ok {
		glog.InfoDepth(depth, fmt.Sprintf(format, a...))
	}
}

func (e *GlogLogger) Debug(a ...interface{}) {
	if ok := CheckLevel(DebugLevel); ok {
		glog.InfoDepth(depth, getLevel()+" ", e.prefix, fmt.Sprint(a...))
	}
}
func (e *GlogLogger) Debugf(format string, a ...interface{}) {
	if ok := CheckLevel(DebugLevel); ok {
		glog.InfoDepth(depth, fmt.Sprintf(getLevel()+" "+format, a...))
	}
}

func (e *GlogLogger) Error(a ...interface{}) {
	if ok := CheckLevel(ErrorLevel); ok {
		glog.ErrorDepth(depth, a...)
	}
}
func (e *GlogLogger) Errorf(format string, a ...interface{}) {
	if ok := CheckLevel(ErrorLevel); ok {
		glog.ErrorDepth(depth, fmt.Sprintf(format, a...))
	}
}

func (e *GlogLogger) Fatal(a ...interface{}) {
	if ok := CheckLevel(FatalLevel); ok {
		glog.FatalDepth(depth, a...)
	}
}
