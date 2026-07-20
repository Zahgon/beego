package session

import (
	"context"
	"net/http"
	"os"
	"sync"
)

var (
	filepder      = &FileProvider{}
	gcmaxlifetime int64
)

type FileSessionStore struct {
	sid    string
	lock   sync.RWMutex
	values map[interface{}]interface{}
}

func (fs *FileSessionStore) Set(ctx context.Context, key, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (fs *FileSessionStore) Get(ctx context.Context, key interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (fs *FileSessionStore) Delete(ctx context.Context, key interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (fs *FileSessionStore) Flush(context.Context) error { _ = "STUB: not implemented"; return nil }

func (fs *FileSessionStore) SessionID(context.Context) string { _ = "STUB: not implemented"; return "" }

func (fs *FileSessionStore) SessionRelease(ctx context.Context, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func (fs *FileSessionStore) SessionReleaseIfPresent(ctx context.Context, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func (fs *FileSessionStore) releaseSession(_ context.Context, _ http.ResponseWriter, createIfNotExist bool) {
	_ = "STUB: not implemented"
	return
}

type FileProvider struct {
	lock        sync.RWMutex
	maxlifetime int64
	savePath    string
}

func (fp *FileProvider) SessionInit(ctx context.Context, maxlifetime int64, savePath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (fp *FileProvider) SessionRead(ctx context.Context, sid string) (Store, error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}

func (fp *FileProvider) SessionExist(ctx context.Context, sid string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (fp *FileProvider) SessionDestroy(ctx context.Context, sid string) error {
	_ = "STUB: not implemented"
	return nil
}

func (fp *FileProvider) SessionGC(context.Context) { _ = "STUB: not implemented"; return }

func (fp *FileProvider) SessionAll(context.Context) int { _ = "STUB: not implemented"; return 0 }

func (fp *FileProvider) SessionRegenerate(ctx context.Context, oldsid, sid string) (Store, error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}

func gcpath(path string, info os.FileInfo, err error) error { _ = "STUB: not implemented"; return nil }

type activeSession struct {
	total int
}

func (as *activeSession) visit(paths string, f os.FileInfo, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	Register("file", filepder)
}
