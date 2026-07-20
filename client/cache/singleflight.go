package cache

import (
	"context"
	"time"

	"golang.org/x/sync/singleflight"
)

type SingleflightCache struct {
	Cache
	group      *singleflight.Group
	expiration time.Duration
	loadFunc   func(ctx context.Context, key string) (any, error)
}

func NewSingleflightCache(c Cache, expiration time.Duration,
	loadFunc func(ctx context.Context, key string) (any, error),
) (Cache, error) {
	_ = "STUB: not implemented"
	return *new(Cache), nil
}

func (s *SingleflightCache) Get(ctx context.Context, key string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
