package alils

import (
	"sync"

	"github.com/beego/beego/v2/core/logs"
)

const (
	CacheSize int = 64

	Delimiter string = "##"
)

type Config struct {
	Project   string   `json:"project"`
	Endpoint  string   `json:"endpoint"`
	KeyID     string   `json:"key_id"`
	KeySecret string   `json:"key_secret"`
	LogStore  string   `json:"log_store"`
	Topics    []string `json:"topics"`
	Source    string   `json:"source"`
	Level     int      `json:"level"`
	FlushWhen int      `json:"flush_when"`
	Formatter string   `json:"formatter"`
}

type aliLSWriter struct {
	store    *LogStore
	group    []*LogGroup
	withMap  bool
	groupMap map[string]*LogGroup
	lock     *sync.Mutex
	Config
	formatter logs.LogFormatter
}

func NewAliLS() logs.Logger { _ = "STUB: not implemented"; return *new(logs.Logger) }

func (c *aliLSWriter) Init(config string) error { _ = "STUB: not implemented"; return nil }

func (c *aliLSWriter) Format(lm *logs.LogMsg) string { _ = "STUB: not implemented"; return "" }

func (c *aliLSWriter) SetFormatter(f logs.LogFormatter) { _ = "STUB: not implemented"; return }

func (c *aliLSWriter) WriteMsg(lm *logs.LogMsg) error { _ = "STUB: not implemented"; return nil }

func (c *aliLSWriter) Flush() { _ = "STUB: not implemented"; return }

func (c *aliLSWriter) Destroy() { _ = "STUB: not implemented"; return }

func (c *aliLSWriter) flush(lg *LogGroup) { _ = "STUB: not implemented"; return }

func init() {
	logs.Register(logs.AdapterAliLS, NewAliLS)
}
