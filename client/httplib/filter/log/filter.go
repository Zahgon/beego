package log

import (
	"io"

	"github.com/beego/beego/v2/client/httplib"
)

type FilterChainBuilder struct {
	printableContentTypes []string
	log                   func(f interface{}, v ...interface{})
}

type BuilderOption func(*FilterChainBuilder)

type logInfo struct {
	req  []byte
	resp []byte
	err  error
}

var defaultprintableContentTypes = []string{
	"text/plain", "text/xml", "text/html", "text/csv",
	"text/calendar", "text/javascript", "text/javascript",
	"text/css",
}

func NewFilterChainBuilder(opts ...BuilderOption) *FilterChainBuilder {
	_ = "STUB: not implemented"
	return nil
}

func WithLog(f func(f interface{}, v ...interface{})) BuilderOption {
	_ = "STUB: not implemented"
	return *new(BuilderOption)
}

func WithprintableContentTypes(types []string) BuilderOption {
	_ = "STUB: not implemented"
	return *new(BuilderOption)
}

func (builder *FilterChainBuilder) FilterChain(next httplib.Filter) httplib.Filter {
	_ = "STUB: not implemented"
	return *new(httplib.Filter)
}

func (builder *FilterChainBuilder) shouldPrintBody(contentType string, body io.ReadCloser) bool {
	_ = "STUB: not implemented"
	return false
}

func contains(s []string, e string) bool { _ = "STUB: not implemented"; return false }

func (info *logInfo) print(log func(f interface{}, v ...interface{})) {
	_ = "STUB: not implemented"
	return
}
