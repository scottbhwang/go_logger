package logger

import (
	"bytes"
	"context"
	"strings"
)

// https://pkg.go.dev/google.golang.org/grpc/metadata#Join

type prefixKey struct{}

type Prefix struct {
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

func getPrefixFromContext(ctx context.Context) Prefix {
	p, ok := fromContext(ctx)
	if !ok {
		return newPrefix()
	}
	return p
}

func (e Prefix) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for k, v := range e.Data {
		if buffer.Len() > 1 {
			buffer.WriteString(" ")
		}
		buffer.WriteString(k + ":" + v)
	}
	buffer.WriteString("]")
	return buffer.String()
}

func (e Prefix) Append(key, value string) Prefix {
	lowKey := strings.ToLower(key)
	e.Data[lowKey] = value
	return e
}
