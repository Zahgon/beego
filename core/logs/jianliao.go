package logs

type JLWriter struct {
	AuthorName  string `json:"authorname"`
	Title       string `json:"title"`
	WebhookURL  string `json:"webhookurl"`
	RedirectURL string `json:"redirecturl,omitempty"`
	ImageURL    string `json:"imageurl,omitempty"`
	Level       int    `json:"level"`

	formatter LogFormatter
	Formatter string `json:"formatter"`
}

func newJLWriter() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (s *JLWriter) Init(config string) error { _ = "STUB: not implemented"; return nil }

func (s *JLWriter) Format(lm *LogMsg) string { _ = "STUB: not implemented"; return "" }

func (s *JLWriter) SetFormatter(f LogFormatter) { _ = "STUB: not implemented"; return }

func (s *JLWriter) WriteMsg(lm *LogMsg) error { _ = "STUB: not implemented"; return nil }

func (s *JLWriter) Flush() { _ = "STUB: not implemented"; return }

func (s *JLWriter) Destroy() { _ = "STUB: not implemented"; return }

func init() {
	Register(AdapterJianLiao, newJLWriter)
}
