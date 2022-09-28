package logger

import (
	"context"
	"strings"
)

// https://pkg.go.dev/google.golang.org/grpc/metadata#Join

type prefixKey struct{}
type Prefix struct {
	Str  string
	Data map[string]string
}

func newPrefix() Prefix {
	return Prefix{
		Data: make(map[string]string),
	}
}

func fromContext(ctx context.Context) (Prefix, bool) {
	p, ok := ctx.Value(prefixKey{}).(Prefix)
	if !ok {
		return newPrefix(), false
	}
	return p, true
}

func (e Prefix) Copy() Prefix {
	copyP := newPrefix()
	copyP.Str = e.Str
	for k, v := range e.Data {
		copyP.Data[k] = v
	}
	return copyP
}

func AppendPrefix(ctx context.Context, key, value string) context.Context {
	p, ok := fromContext(ctx)
	if ok {
		v := p.Copy().Append(strings.ToLower(key), value)
		ctx = context.WithValue(ctx, prefixKey{}, v)
		return ctx
	} else {
		v := newPrefix().Append(strings.ToLower(key), value)
		ctx = context.WithValue(ctx, prefixKey{}, v)
		return ctx
	}
}

func getPrefixStrFromContext(ctx context.Context) string {
	p, ok := fromContext(ctx)
	if !ok {
		return ""
	}
	return p.String()
}

func (e Prefix) String() string {
	if len(e.Str) > 0 {
		return e.Str + " "
	}
	return ""
}

func (e Prefix) Append(key, value string) Prefix {
	lowKey := strings.ToLower(key)
	e.Data[lowKey] = value
	if len(e.Str) > 0 {
		e.Str += ","
	}
	e.Str += lowKey + "=" + value
	return e
}
