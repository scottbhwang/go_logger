package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
)

var log ILogger
var Level = InfoLevel
var Path string

func init() {
	flag.Var(&Level, "logger_level", "日志等级 debug,info,error")
	flag.StringVar(&Path, "logger_path", "./", "日志等级 debug,info,error")
	log = &GlogLogger{prefix: func(level LogLevel) string {
		return DefaultPrefix(level) + " "
	}}
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", Path)
}

func WithContext(ctx context.Context) ILogger {
	prefix := getPrefixFromContext(ctx)
	return &GlogLogger{prefix: func(level LogLevel) string {
		return DefaultPrefix(level) + " " + prefix.String() + " "
	}}
}

func DefaultPrefix(level LogLevel) string {
	return level.String()
}

func Info(a ...interface{}) {
	log.Info(a...)
}
func Infof(format string, a ...interface{}) {
	log.Infof(format, a...)
}

func Debug(a ...interface{}) {
	log.Debug(a...)
}
func Debugf(format string, a ...interface{}) {
	log.Debugf(format, a...)
}

func Error(a ...interface{}) {
	log.Error(a...)
}
func Errorf(format string, a ...interface{}) {
	log.Errorf(format, a...)
}

func Fatal(a ...interface{}) {
	log.Fatal(a...)
}

func SprintPretty(data interface{}) string {
	val, _ := json.MarshalIndent(data, "", "    ")
	return string(val)
}

func Sprint(data interface{}) string {
	val, _ := json.Marshal(data)
	return string(val)
}

func SprintJsonStringPretty(j string) string {
	var buf bytes.Buffer
	_ = json.Indent(&buf, []byte(j), "", "    ")
	return buf.String()
}

func CheckLevel(level LogLevel) bool {
	return Level <= level
}
