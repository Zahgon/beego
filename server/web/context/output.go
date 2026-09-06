package context

import (
	"strings"

	"google.golang.org/protobuf/proto"
)

type BeegoOutput struct {
	Context    *Context
	Status     int
	EnableGzip bool
}

func NewOutput() *BeegoOutput { _ = "STUB: not implemented"; return nil }

func (output *BeegoOutput) Reset(ctx *Context) { _ = "STUB: not implemented"; return }

func (output *BeegoOutput) Header(key, val string) { _ = "STUB: not implemented"; return }

func (output *BeegoOutput) Body(content []byte) error { _ = "STUB: not implemented"; return nil }

func (output *BeegoOutput) Cookie(name string, value string, others ...interface{}) {
	_ = "STUB: not implemented"
	return
}

var cookieNameSanitizer = strings.NewReplacer("\n", "-", "\r", "-")

func sanitizeName(n string) string { _ = "STUB: not implemented"; return "" }

var cookieValueSanitizer = strings.NewReplacer("\n", " ", "\r", " ", ";", " ")

func sanitizeValue(v string) string { _ = "STUB: not implemented"; return "" }

func jsonRenderer(value interface{}) Renderer { _ = "STUB: not implemented"; return *new(Renderer) }

func errorRenderer(err error) Renderer { _ = "STUB: not implemented"; return *new(Renderer) }

func (output *BeegoOutput) JSON(data interface{}, hasIndent bool, encoding bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (output *BeegoOutput) YAML(data interface{}) error { _ = "STUB: not implemented"; return nil }

func (output *BeegoOutput) Proto(data proto.Message) error { _ = "STUB: not implemented"; return nil }

func (output *BeegoOutput) JSONP(data interface{}, hasIndent bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (output *BeegoOutput) XML(data interface{}, hasIndent bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (output *BeegoOutput) ServeFormatted(data interface{}, hasIndent bool, hasEncode ...bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (output *BeegoOutput) Download(file string, filename ...string) {
	_ = "STUB: not implemented"
	return
}

func (output *BeegoOutput) ContentType(ext string) { _ = "STUB: not implemented"; return }

func (output *BeegoOutput) SetStatus(status int) { _ = "STUB: not implemented"; return }

func (output *BeegoOutput) IsCachable() bool { _ = "STUB: not implemented"; return false }

func (output *BeegoOutput) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (output *BeegoOutput) IsOk() bool { _ = "STUB: not implemented"; return false }

func (output *BeegoOutput) IsSuccessful() bool { _ = "STUB: not implemented"; return false }

func (output *BeegoOutput) IsRedirect() bool { _ = "STUB: not implemented"; return false }

func (output *BeegoOutput) IsForbidden() bool { _ = "STUB: not implemented"; return false }

func (output *BeegoOutput) IsNotFound() bool { _ = "STUB: not implemented"; return false }

func (output *BeegoOutput) IsClientError() bool { _ = "STUB: not implemented"; return false }

func (output *BeegoOutput) IsServerError() bool { _ = "STUB: not implemented"; return false }

func stringsToJSON(str string) string { _ = "STUB: not implemented"; return "" }

func (output *BeegoOutput) Session(name interface{}, value interface{}) {
	_ = "STUB: not implemented"
	return
}
