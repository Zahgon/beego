package utils

import (
	"bytes"
	"io"
	"net/smtp"
	"net/textproto"
	"sync"
)

const (
	maxLineLength = 76

	upperhex = "0123456789ABCDEF"
)

type Email struct {
	Auth        smtp.Auth
	Identity    string `json:"identity"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	From        string `json:"from"`
	To          []string
	Bcc         []string
	Cc          []string
	Subject     string
	Text        string
	HTML        string
	Headers     textproto.MIMEHeader
	Attachments []*Attachment
	ReadReceipt []string
}

type Attachment struct {
	Filename string
	Header   textproto.MIMEHeader
	Content  []byte
}

func NewEMail(config string) *Email { _ = "STUB: not implemented"; return nil }

func (e *Email) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *Email) AttachFile(args ...string) (a *Attachment, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Email) Attach(r io.Reader, filename string, args ...string) (a *Attachment, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Email) Send() error { _ = "STUB: not implemented"; return nil }

func quotePrintEncode(w io.Writer, s string) error { _ = "STUB: not implemented"; return nil }

func isPrintable(c byte) bool { _ = "STUB: not implemented"; return false }

func qpEscape(dest []byte, c byte) { _ = "STUB: not implemented"; return }

func headerToBytes(w io.Writer, t textproto.MIMEHeader) error {
	_ = "STUB: not implemented"
	return nil
}

func base64Wrap(w io.Writer, b []byte) { _ = "STUB: not implemented"; return }

func qEncode(charset, s string) string { _ = "STUB: not implemented"; return "" }

func needsEncoding(s string) bool { _ = "STUB: not implemented"; return false }

func encodeWord(charset, s string) string { _ = "STUB: not implemented"; return "" }

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func getBuffer() *bytes.Buffer { _ = "STUB: not implemented"; return nil }

func putBuffer(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }
