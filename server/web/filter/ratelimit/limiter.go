package ratelimit

import (
	"sync"
	"time"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

type limiterOption func(l *limiter)

type limiter struct {
	sync.RWMutex
	capacity      uint
	rate          time.Duration
	buckets       map[string]bucket
	bucketFactory func(opts ...bucketOption) bucket
	sessionKey    func(ctx *context.Context) string
	resp          RejectionResponse
}

type RejectionResponse struct {
	code int
	body string
}

const perRequestConsumedAmount = 1

var defaultRejectionResponse = RejectionResponse{
	code: 429,
	body: "too many requests",
}

func NewLimiter(opts ...limiterOption) web.FilterFunc {
	_ = "STUB: not implemented"
	return *new(web.FilterFunc)
}

func WithSessionKey(f func(ctx *context.Context) string) limiterOption {
	_ = "STUB: not implemented"
	return *new(limiterOption)
}

func WithRate(r time.Duration) limiterOption { _ = "STUB: not implemented"; return *new(limiterOption) }

func WithCapacity(c uint) limiterOption { _ = "STUB: not implemented"; return *new(limiterOption) }

func WithBucketFactory(f func(opts ...bucketOption) bucket) limiterOption {
	_ = "STUB: not implemented"
	return *new(limiterOption)
}

func WithRejectionResponse(resp RejectionResponse) limiterOption {
	_ = "STUB: not implemented"
	return *new(limiterOption)
}

func (l *limiter) take(amount uint, ctx *context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (l *limiter) getBucket(ctx *context.Context) bucket {
	_ = "STUB: not implemented"
	return *new(bucket)
}

func (l *limiter) createBucket(key string) bucket { _ = "STUB: not implemented"; return *new(bucket) }

func defaultSessionKey(ctx *context.Context) string { _ = "STUB: not implemented"; return "" }

func RemoteIPSessionKey(ctx *context.Context) string { _ = "STUB: not implemented"; return "" }
