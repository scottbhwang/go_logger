package logger

import (
	"bytes"
	"fmt"
	"github.com/petermattis/goid"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type FmtLogger struct{}

func (e *FmtLogger) Info(a ...interface{}) {
	if ok := CheckLevel(InfoLevel); ok {
		fmt.Println(fmtPrefix(InfoLevel), fmt.Sprint(a...))
	}
}
func (e *FmtLogger) Infof(format string, a ...interface{}) {
	if ok := CheckLevel(InfoLevel); ok {
		fmt.Println(fmt.Sprintf(fmtPrefix(InfoLevel)+format, a...))
	}
}

func (e *FmtLogger) Debug(a ...interface{}) {
	if ok := CheckLevel(DebugLevel); ok {
		fmt.Println(fmtPrefix(DebugLevel), fmt.Sprint(a...))
	}
}
func (e *FmtLogger) Debugf(format string, a ...interface{}) {
	if ok := CheckLevel(DebugLevel); ok {
		fmt.Println(fmt.Sprintf(fmtPrefix(DebugLevel)+format, a...))
	}
}

func (e *FmtLogger) Error(a ...interface{}) {
	if ok := CheckLevel(ErrorLevel); ok {
		_, _ = os.Stderr.WriteString(fmtPrefix(ErrorLevel) + fmt.Sprintln(a...))
	}
}
func (e *FmtLogger) Errorf(format string, a ...interface{}) {
	if ok := CheckLevel(ErrorLevel); ok {
		_, _ = os.Stderr.WriteString(fmt.Sprintf(fmtPrefix(ErrorLevel)+format, a...))
	}
}

func (e *FmtLogger) Fatal(a ...interface{}) {
	if ok := CheckLevel(FatalLevel); ok {
		stackSlice := make([]byte, 2048)
		s := runtime.Stack(stackSlice, false)
		ss := bytes.SplitN(stackSlice[0:s], []byte("\n"), 3)[2]
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s\n%s", fmt.Sprint(a...), ss))
	}
}

var fmtPrefix = func(level LogLevel) string {
	return fmt.Sprintf("%s [%s] %s %s", getTime(), level, getPid(), getCaller())
}

var getPid = func() string {
	return strconv.FormatInt(goid.Get(), 10)
}
var getTime = func() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}

var getCaller = func() string {
	_, file, line, ok := runtime.Caller(4)
	if !ok {
		file = "???"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}

	return file + ":" + strconv.Itoa(line)
}
