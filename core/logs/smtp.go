package logs

import (
	"net/smtp"
)

type SMTPWriter struct {
	Username           string   `json:"username"`
	Password           string   `json:"password"`
	Host               string   `json:"host"`
	Subject            string   `json:"subject"`
	FromAddress        string   `json:"fromAddress"`
	RecipientAddresses []string `json:"sendTos"`
	Level              int      `json:"level"`

	InsecureSkipVerify bool `json:"insecureSkipVerify"`

	formatter LogFormatter
	Formatter string `json:"formatter"`
}

func newSMTPWriter() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (s *SMTPWriter) Init(config string) error { _ = "STUB: not implemented"; return nil }

func (s *SMTPWriter) getSMTPAuth(host string) smtp.Auth {
	_ = "STUB: not implemented"
	return *new(smtp.Auth)
}

func (s *SMTPWriter) SetFormatter(f LogFormatter) { _ = "STUB: not implemented"; return }

func (s *SMTPWriter) sendMail(hostAddressWithPort string, auth smtp.Auth, fromAddress string, recipients []string, msgContent []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SMTPWriter) Format(lm *LogMsg) string { _ = "STUB: not implemented"; return "" }

func (s *SMTPWriter) WriteMsg(lm *LogMsg) error { _ = "STUB: not implemented"; return nil }

func (s *SMTPWriter) Flush() { _ = "STUB: not implemented"; return }

func (s *SMTPWriter) Destroy() { _ = "STUB: not implemented"; return }

func init() {
	Register(AdapterMail, newSMTPWriter)
}
