package context

import (
	"bufio"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/beego/beego/v2/server/web/session"
)

const (
	ApplicationJSON  = "application/json"
	ApplicationXML   = "application/xml"
	ApplicationForm  = "application/x-www-form-urlencoded"
	ApplicationProto = "application/x-protobuf"
	ApplicationYAML  = "application/x-yaml"
	TextXML          = "text/xml"

	formatTime      = "15:04:05"
	formatDate      = "2006-01-02"
	formatDateTime  = "2006-01-02 15:04:05"
	formatDateTimeT = "2006-01-02T15:04:05"
)

func NewContext() *Context { _ = "STUB: not implemented"; return nil }

type Context struct {
	Input          *BeegoInput
	Output         *BeegoOutput
	Request        *http.Request
	ResponseWriter *Response
	_xsrfToken     string
}

func (ctx *Context) Bind(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (ctx *Context) Resp(data interface{}) error { _ = "STUB: not implemented"; return nil }

func (ctx *Context) JSONResp(data interface{}) error { _ = "STUB: not implemented"; return nil }

func (ctx *Context) XMLResp(data interface{}) error { _ = "STUB: not implemented"; return nil }

func (ctx *Context) YamlResp(data interface{}) error { _ = "STUB: not implemented"; return nil }

func (ctx *Context) ProtoResp(data proto.Message) error { _ = "STUB: not implemented"; return nil }

func (ctx *Context) BindYAML(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (ctx *Context) BindForm(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (ctx *Context) BindJSON(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (ctx *Context) BindProtobuf(obj proto.Message) error { _ = "STUB: not implemented"; return nil }

func (ctx *Context) BindXML(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func ParseForm(form url.Values, obj interface{}) error { _ = "STUB: not implemented"; return nil }

func isStructPtr(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func (ctx *Context) Reset(rw http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (ctx *Context) Redirect(status int, localurl string) { _ = "STUB: not implemented"; return }

func (ctx *Context) Abort(status int, body string) { _ = "STUB: not implemented"; return }

func (ctx *Context) WriteString(content string) { _ = "STUB: not implemented"; return }

func (ctx *Context) GetCookie(key string) string { _ = "STUB: not implemented"; return "" }

func (ctx *Context) SetCookie(name string, value string, others ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ctx *Context) GetSecureCookie(Secret, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (ctx *Context) SetSecureCookie(Secret, name, value string, others ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ctx *Context) XSRFToken(key string, expire int64) string {
	_ = "STUB: not implemented"
	return ""
}

func (ctx *Context) CheckXSRFCookie() bool { _ = "STUB: not implemented"; return false }

func (ctx *Context) RenderMethodResult(result interface{}) { _ = "STUB: not implemented"; return }

func (ctx *Context) Session() (store session.Store, err error) {
	_ = "STUB: not implemented"
	return *new(session.Store), nil
}

type Response struct {
	http.ResponseWriter
	Started bool
	Status  int
	Elapsed time.Duration
}

func (r *Response) reset(rw http.ResponseWriter) { _ = "STUB: not implemented"; return }

func (r *Response) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *Response) WriteHeader(code int) { _ = "STUB: not implemented"; return }

func (r *Response) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

func (r *Response) Flush() { _ = "STUB: not implemented"; return }

func (r *Response) CloseNotify() <-chan bool { _ = "STUB: not implemented"; return nil }

func (r *Response) Pusher() (pusher http.Pusher) {
	_ = "STUB: not implemented"
	return *new(http.Pusher)
}
