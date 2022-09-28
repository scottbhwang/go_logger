package logger

import (
	"context"
	"encoding/json"
	"flag"
)

var levelFlag string

var log ILogger
var Level = InfoLevel

func init() {
	flag.Var(&Level, "logger_level", "日志等级 debug,info,error")
	log = &GlogLogger{}
	flag.Set("alsologtostderr", "true")
}

func WithContext(ctx context.Context) ILogger {
	prefix := getPrefixStrFromContext(ctx)
	return &GlogLogger{prefix: prefix}
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

func CheckLevel(level LogLevel) bool {
	return Level <= level
}
