package cache

import (
	"context"
	"time"
)

type readThroughCache struct {
	Cache
	expiration time.Duration
	loadFunc   func(ctx context.Context, key string) (any, error)
}

func NewReadThroughCache(cache Cache, expiration time.Duration,
	loadFunc func(ctx context.Context, key string) (any, error),
) (Cache, error) {
	_ = "STUB: not implemented"
	return *new(Cache), nil
}

func (c *readThroughCache) Get(ctx context.Context, key string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
