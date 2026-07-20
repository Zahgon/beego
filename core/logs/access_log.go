package logs

import (
	"time"
)

const (
	apacheFormatPattern = "%s - - [%s] \"%s %d %d\" %f %s %s"
	apacheFormat        = "APACHE_FORMAT"
	jsonFormat          = "JSON_FORMAT"
)

type AccessLogRecord struct {
	RemoteAddr     string        `json:"remote_addr"`
	RequestTime    time.Time     `json:"request_time"`
	RequestMethod  string        `json:"request_method"`
	Request        string        `json:"request"`
	ServerProtocol string        `json:"server_protocol"`
	Host           string        `json:"host"`
	Status         int           `json:"status"`
	BodyBytesSent  int64         `json:"body_bytes_sent"`
	ElapsedTime    time.Duration `json:"elapsed_time"`
	HTTPReferrer   string        `json:"http_referrer"`
	HTTPUserAgent  string        `json:"http_user_agent"`
	RemoteUser     string        `json:"remote_user"`
}

func (r *AccessLogRecord) json() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func disableEscapeHTML(i interface{}) { _ = "STUB: not implemented"; return }

func AccessLog(r *AccessLogRecord, format string) { _ = "STUB: not implemented"; return }

func (r *AccessLogRecord) format(format string) string { _ = "STUB: not implemented"; return "" }
