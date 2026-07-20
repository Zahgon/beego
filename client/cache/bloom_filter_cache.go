package cache

import (
	"context"
	"time"
)

type BloomFilterCache struct {
	Cache
	blm        BloomFilter
	loadFunc   func(ctx context.Context, key string) (any, error)
	expiration time.Duration
}

type BloomFilter interface {
	Test(data string) bool
	Add(data string)
}

func NewBloomFilterCache(cache Cache, ln func(ctx context.Context, key string) (any, error), blm BloomFilter,
	expiration time.Duration,
) (*BloomFilterCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bfc *BloomFilterCache) Get(ctx context.Context, key string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
