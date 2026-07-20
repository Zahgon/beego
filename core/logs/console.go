package logs

type brush func(string) string

func newBrush(color string) brush { _ = "STUB: not implemented"; return *new(brush) }

var colors = []brush{
	newBrush("1;37"),
	newBrush("1;36"),
	newBrush("1;35"),
	newBrush("1;31"),
	newBrush("1;33"),
	newBrush("1;32"),
	newBrush("1;34"),
	newBrush("1;44"),
}

type consoleWriter struct {
	lg        *logWriter
	formatter LogFormatter
	Formatter string `json:"formatter"`
	Level     int    `json:"level"`
	Colorful  bool   `json:"color"`
}

func (c *consoleWriter) Format(lm *LogMsg) string { _ = "STUB: not implemented"; return "" }

func (c *consoleWriter) SetFormatter(f LogFormatter) { _ = "STUB: not implemented"; return }

func NewConsole() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func newConsole() *consoleWriter { _ = "STUB: not implemented"; return nil }

func (c *consoleWriter) Init(config string) error { _ = "STUB: not implemented"; return nil }

func (c *consoleWriter) WriteMsg(lm *LogMsg) error { _ = "STUB: not implemented"; return nil }

func (c *consoleWriter) Destroy() { _ = "STUB: not implemented"; return }

func (c *consoleWriter) Flush() { _ = "STUB: not implemented"; return }

func init() {
	Register(AdapterConsole, NewConsole)
}
