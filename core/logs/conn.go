package logs

import (
	"io"
)

type connWriter struct {
	lg             *logWriter
	innerWriter    io.WriteCloser
	formatter      LogFormatter
	Formatter      string `json:"formatter"`
	ReconnectOnMsg bool   `json:"reconnectOnMsg"`
	Reconnect      bool   `json:"reconnect"`
	Net            string `json:"net"`
	Addr           string `json:"addr"`
	Level          int    `json:"level"`
}

func NewConn() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (c *connWriter) Format(lm *LogMsg) string { _ = "STUB: not implemented"; return "" }

func (c *connWriter) Init(config string) error { _ = "STUB: not implemented"; return nil }

func (c *connWriter) SetFormatter(f LogFormatter) { _ = "STUB: not implemented"; return }

func (c *connWriter) WriteMsg(lm *LogMsg) error { _ = "STUB: not implemented"; return nil }

func (c *connWriter) Flush() { _ = "STUB: not implemented"; return }

func (c *connWriter) Destroy() { _ = "STUB: not implemented"; return }

func (c *connWriter) connect() error { _ = "STUB: not implemented"; return nil }

func (c *connWriter) needToConnectOnMsg() bool { _ = "STUB: not implemented"; return false }

func init() {
	Register(AdapterConn, NewConn)
}
