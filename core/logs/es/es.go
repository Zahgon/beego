package es

import (
	"github.com/elastic/go-elasticsearch/v6"

	"github.com/beego/beego/v2/core/logs"
)

func NewES() logs.Logger { _ = "STUB: not implemented"; return *new(logs.Logger) }

type esLogger struct {
	*elasticsearch.Client
	DSN       string `json:"dsn"`
	Level     int    `json:"level"`
	formatter logs.LogFormatter
	Formatter string `json:"formatter"`

	indexNaming IndexNaming
}

func (el *esLogger) Format(lm *logs.LogMsg) string { _ = "STUB: not implemented"; return "" }

func (el *esLogger) SetFormatter(f logs.LogFormatter) { _ = "STUB: not implemented"; return }

func (el *esLogger) Init(config string) error { _ = "STUB: not implemented"; return nil }

func (el *esLogger) WriteMsg(lm *logs.LogMsg) error { _ = "STUB: not implemented"; return nil }

func (el *esLogger) Destroy() { _ = "STUB: not implemented"; return }

func (el *esLogger) Flush() { _ = "STUB: not implemented"; return }

type LogDocument struct {
	Timestamp string `json:"timestamp"`
	Msg       string `json:"msg"`
}

func init() {
	logs.Register(logs.AdapterEs, NewES)
}
