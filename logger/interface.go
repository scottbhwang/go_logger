package logger

type ILogger interface {
	Info(a ...interface{})
	Infof(format string, a ...interface{})
	//InfoWithCtx(ctx context.Context, a ...interface{})
	Debug(a ...interface{})
	Debugf(format string, a ...interface{})
	//DebugWithCtx(ctx context.Context, a ...interface{})
	Error(a ...interface{})
	Errorf(format string, a ...interface{})
	Fatal(a ...interface{})
}
