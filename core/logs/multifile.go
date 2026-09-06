package logs

type multiFileLogWriter struct {
	writers       [LevelDebug + 1 + 1]*fileLogWriter
	fullLogWriter *fileLogWriter
	Separate      []string `json:"separate"`
}

var levelNames = [...]string{"emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"}

func (f *multiFileLogWriter) Init(config string) error { _ = "STUB: not implemented"; return nil }

func (*multiFileLogWriter) Format(lm *LogMsg) string { _ = "STUB: not implemented"; return "" }

func (f *multiFileLogWriter) SetFormatter(fmt LogFormatter) { _ = "STUB: not implemented"; return }

func (f *multiFileLogWriter) Destroy() { _ = "STUB: not implemented"; return }

func (f *multiFileLogWriter) WriteMsg(lm *LogMsg) error { _ = "STUB: not implemented"; return nil }

func (f *multiFileLogWriter) Flush() { _ = "STUB: not implemented"; return }

func newFilesWriter() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func init() {
	Register(AdapterMultiFile, newFilesWriter)
}
