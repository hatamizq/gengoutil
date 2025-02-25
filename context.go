package gengoutil

import (
	"context"
	"time"
)

// new implementation for context without deadline; used for background process but need context from request
type ContextWithoutDeadline struct {
	Ctx context.Context
}

func (*ContextWithoutDeadline) Deadline() (time.Time, bool) { return time.Time{}, false }
func (*ContextWithoutDeadline) Done() <-chan struct{}       { return nil }
func (*ContextWithoutDeadline) Err() error                  { return nil }

func (c *ContextWithoutDeadline) Value(key interface{}) interface{} {
	return c.Ctx.Value(key)
}

func NewContextWithoutDeadline(ctx context.Context) context.Context {
	return &ContextWithoutDeadline{ctx}
}

func GetContextValue[T any](ctx context.Context, key string) T {
	value := ctx.Value(key)

	if value == nil {
		return ZeroOf[T]()
	}

	parsedValue, ok := value.(T)
	if !ok {
		return ZeroOf[T]()
	}
	return parsedValue
}
