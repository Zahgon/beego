package web

import (
	"bytes"
	"errors"
	"os"
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru"

	"github.com/beego/beego/v2/server/web/context"
)

var errNotStaticRequest = errors.New("request not a static file request")

func serverStaticRouter(ctx *context.Context) { _ = "STUB: not implemented"; return }

type serveContentHolder struct {
	data       []byte
	modTime    time.Time
	size       int64
	originSize int64
	encoding   string
}

type serveContentReader struct {
	*bytes.Reader
}

var (
	staticFileLruCache *lru.Cache
	lruLock            sync.RWMutex
)

func openFile(filePath string, fi os.FileInfo, acceptEncoding string) (bool, string, *serveContentHolder, *serveContentReader, error) {
	_ = "STUB: not implemented"
	return false, "", nil, nil, nil
}

func isOk(s *serveContentHolder, fi os.FileInfo) bool { _ = "STUB: not implemented"; return false }

func isStaticCompress(filePath string) bool { _ = "STUB: not implemented"; return false }

func searchFile(ctx *context.Context) (string, os.FileInfo, error) {
	_ = "STUB: not implemented"
	return "", *new(os.FileInfo), nil
}

func lookupFile(ctx *context.Context) (bool, string, os.FileInfo, error) {
	_ = "STUB: not implemented"
	return false, "", *new(os.FileInfo), nil
}
