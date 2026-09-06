package cache

import (
	"context"
	"time"
)

type FileCacheItem struct {
	Data       interface{}
	Lastaccess time.Time
	Expired    time.Time
}

var (
	FileCachePath           = "cache"
	FileCacheFileSuffix     = ".bin"
	FileCacheDirectoryLevel = 2
	FileCacheEmbedExpiry    time.Duration
)

type FileCache struct {
	CachePath      string
	FileSuffix     string
	DirectoryLevel int
	EmbedExpiry    int
}

func NewFileCache() Cache { _ = "STUB: not implemented"; return *new(Cache) }

func (fc *FileCache) StartAndGC(config string) error { _ = "STUB: not implemented"; return nil }

func (fc *FileCache) Init() error { _ = "STUB: not implemented"; return nil }

func (fc *FileCache) getCacheFileName(key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (fc *FileCache) Get(ctx context.Context, key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fc *FileCache) GetMulti(ctx context.Context, keys []string) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fc *FileCache) Put(ctx context.Context, key string, val interface{}, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (fc *FileCache) Delete(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (fc *FileCache) Incr(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (fc *FileCache) Decr(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (fc *FileCache) IsExist(ctx context.Context, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (fc *FileCache) ClearAll(context.Context) error { _ = "STUB: not implemented"; return nil }

func exists(path string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func FileGetContents(filename string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func FilePutContents(filename string, content []byte) error { _ = "STUB: not implemented"; return nil }

func GobEncode(data interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func GobDecode(data []byte, to *FileCacheItem) error { _ = "STUB: not implemented"; return nil }

func init() {
	Register("file", NewFileCache)
}
