package logs

type SLACKWriter struct {
	WebhookURL string `json:"webhookurl"`
	Level      int    `json:"level"`
	formatter  LogFormatter
	Formatter  string `json:"formatter"`
}

func newSLACKWriter() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (s *SLACKWriter) Format(lm *LogMsg) string { _ = "STUB: not implemented"; return "" }

func (s *SLACKWriter) SetFormatter(f LogFormatter) { _ = "STUB: not implemented"; return }

func (s *SLACKWriter) Init(config string) error { _ = "STUB: not implemented"; return nil }

func (s *SLACKWriter) WriteMsg(lm *LogMsg) error { _ = "STUB: not implemented"; return nil }

func (s *SLACKWriter) Flush() { _ = "STUB: not implemented"; return }

func (s *SLACKWriter) Destroy() { _ = "STUB: not implemented"; return }

func init() {
	Register(AdapterSlack, newSLACKWriter)
}
