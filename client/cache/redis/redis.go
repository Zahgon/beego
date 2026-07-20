package redis

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/beego/beego/v2/client/cache"
)

const (
	DefaultKey = "beecacheRedis"

	defaultMaxIdle = 3

	defaultTimeout = time.Second * 180
)

type Cache struct {
	p        *redis.Pool
	conninfo string
	dbNum    int

	key      string
	password string
	maxIdle  int

	skipEmptyPrefix bool

	timeout time.Duration
}

func NewRedisCache() cache.Cache { _ = "STUB: not implemented"; return *new(cache.Cache) }

func (rc *Cache) do(commandName string, args ...interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rc *Cache) associate(originKey interface{}) string { _ = "STUB: not implemented"; return "" }

func (rc *Cache) Get(ctx context.Context, key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rc *Cache) GetMulti(ctx context.Context, keys []string) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rc *Cache) Put(ctx context.Context, key string, val interface{}, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (rc *Cache) Delete(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (rc *Cache) IsExist(ctx context.Context, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rc *Cache) Incr(ctx context.Context, key string) error { _ = "STUB: not implemented"; return nil }

func (rc *Cache) Decr(ctx context.Context, key string) error { _ = "STUB: not implemented"; return nil }

func (rc *Cache) ClearAll(context.Context) error { _ = "STUB: not implemented"; return nil }

func (rc *Cache) Scan(pattern string) (keys []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rc *Cache) StartAndGC(config string) error { _ = "STUB: not implemented"; return nil }

func (rc *Cache) parseConf(config string) error { _ = "STUB: not implemented"; return nil }

type redisConfig struct {
	DbNum           string `json:"dbNum"`
	SkipEmptyPrefix string `json:"skipEmptyPrefix"`
	Key             string `json:"key"`

	Conn       string `json:"conn"`
	MaxIdle    string `json:"maxIdle"`
	TimeoutStr string `json:"timeout"`

	dbNum           int
	skipEmptyPrefix bool
	maxIdle         int

	password string

	timeout time.Duration
}

func (cf *redisConfig) parse() error { _ = "STUB: not implemented"; return nil }

func (rc *Cache) connectInit() { _ = "STUB: not implemented"; return }

func init() {
	cache.Register("redis", NewRedisCache)
}
