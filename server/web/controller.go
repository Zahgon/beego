package web

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/url"
	"sync"

	"google.golang.org/protobuf/proto"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/context/param"
	"github.com/beego/beego/v2/server/web/session"
)

var (
	ErrAbort = errors.New("user stop run")

	GlobalControllerRouter = make(map[string][]ControllerComments)
	copyBufferPool         sync.Pool
)

const (
	bytePerKb    = 1024
	copyBufferKb = 32
	filePerm     = 0o666
)

func init() {
	copyBufferPool.New = func() interface{} {
		return make([]byte, bytePerKb*copyBufferKb)
	}
}

type ControllerFilter struct {
	Pattern        string
	Pos            int
	Filter         FilterFunc
	ReturnOnOutput bool
	ResetParams    bool
}

type ControllerFilterComments struct {
	Pattern        string
	Pos            int
	Filter         string
	ReturnOnOutput bool
	ResetParams    bool
}

type ControllerImportComments struct {
	ImportPath  string
	ImportAlias string
}

type ControllerComments struct {
	Method           string
	Router           string
	Filters          []*ControllerFilter
	ImportComments   []*ControllerImportComments
	FilterComments   []*ControllerFilterComments
	AllowHTTPMethods []string
	Params           []map[string]string
	MethodParams     []*param.MethodParam
}

type ControllerCommentsSlice []ControllerComments

func (p ControllerCommentsSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (p ControllerCommentsSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (p ControllerCommentsSlice) Swap(i, j int) { _ = "STUB: not implemented"; return }

type Controller struct {
	Ctx  *context.Context
	Data map[interface{}]interface{}

	controllerName string
	actionName     string
	methodMapping  map[string]func() //method:routertree
	AppController  interface{}

	TplName        string
	ViewPath       string
	Layout         string
	LayoutSections map[string]string
	TplPrefix      string
	TplExt         string
	EnableRender   bool

	EnableXSRF bool
	_xsrfToken string
	XSRFExpire int

	CruSession session.Store
}

type ControllerInterface interface {
	Init(ct *context.Context, controllerName, actionName string, app interface{})
	Prepare()
	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
	Trace()
	Finish()
	Render() error
	XSRFToken() string
	CheckXSRFCookie() bool
	HandlerFunc(fn string) bool
	URLMapping()
}

func (c *Controller) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) Prepare() { _ = "STUB: not implemented"; return }

func (c *Controller) Finish() { _ = "STUB: not implemented"; return }

func (c *Controller) Get() { _ = "STUB: not implemented"; return }

func (c *Controller) Post() { _ = "STUB: not implemented"; return }

func (c *Controller) Delete() { _ = "STUB: not implemented"; return }

func (c *Controller) Put() { _ = "STUB: not implemented"; return }

func (c *Controller) Head() { _ = "STUB: not implemented"; return }

func (c *Controller) Patch() { _ = "STUB: not implemented"; return }

func (c *Controller) Options() { _ = "STUB: not implemented"; return }

func (c *Controller) Trace() { _ = "STUB: not implemented"; return }

func (c *Controller) HandlerFunc(fnname string) bool { _ = "STUB: not implemented"; return false }

func (c *Controller) URLMapping() { _ = "STUB: not implemented"; return }

func (c *Controller) Bind(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (c *Controller) BindYAML(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (c *Controller) BindForm(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (c *Controller) BindJSON(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (c *Controller) BindProtobuf(obj proto.Message) error { _ = "STUB: not implemented"; return nil }

func (c *Controller) BindXML(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (c *Controller) Mapping(method string, fn func()) { _ = "STUB: not implemented"; return }

func (c *Controller) Render() error { _ = "STUB: not implemented"; return nil }

func (c *Controller) RenderString() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *Controller) RenderBytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Controller) renderTemplate() (bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}

func (c *Controller) viewPath() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Redirect(url string, code int) { _ = "STUB: not implemented"; return }

func (c *Controller) SetData(data interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) Abort(code string) { _ = "STUB: not implemented"; return }

func (c *Controller) CustomAbort(status int, body string) { _ = "STUB: not implemented"; return }

func (c *Controller) StopRun() { _ = "STUB: not implemented"; return }

func (c *Controller) URLFor(endpoint string, values ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Controller) JSONResp(data interface{}) error { _ = "STUB: not implemented"; return nil }

func (c *Controller) XMLResp(data interface{}) error { _ = "STUB: not implemented"; return nil }

func (c *Controller) YamlResp(data interface{}) error { _ = "STUB: not implemented"; return nil }

func (c *Controller) Resp(data interface{}) error { _ = "STUB: not implemented"; return nil }

func (c *Controller) ServeJSON(encoding ...bool) error { _ = "STUB: not implemented"; return nil }

func (c *Controller) ServeJSONP() error { _ = "STUB: not implemented"; return nil }

func (c *Controller) ServeXML() error { _ = "STUB: not implemented"; return nil }

func (c *Controller) ServeYAML() error { _ = "STUB: not implemented"; return nil }

func (c *Controller) ServeFormatted(encoding ...bool) error { _ = "STUB: not implemented"; return nil }

func (c *Controller) Input() (url.Values, error) {
	_ = "STUB: not implemented"
	return *new(url.Values), nil
}

func (c *Controller) ParseForm(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (c *Controller) GetString(key string, def ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Controller) GetStrings(key string, def ...[]string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) GetInt(key string, def ...int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Controller) GetInt8(key string, def ...int8) (int8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Controller) GetUint8(key string, def ...uint8) (uint8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Controller) GetInt16(key string, def ...int16) (int16, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Controller) GetUint16(key string, def ...uint16) (uint16, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Controller) GetInt32(key string, def ...int32) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Controller) GetUint32(key string, def ...uint32) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Controller) GetInt64(key string, def ...int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Controller) GetUint64(key string, def ...uint64) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Controller) GetBool(key string, def ...bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Controller) GetFloat(key string, def ...float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Controller) GetFile(key string) (multipart.File, *multipart.FileHeader, error) {
	_ = "STUB: not implemented"
	return *new(multipart.File), nil, nil
}

func (c *Controller) GetFiles(key string) ([]*multipart.FileHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Controller) SaveToFile(fromFile, toFile string) error {
	_ = "STUB: not implemented"
	return nil
}

type onlyWriter struct {
	io.Writer
}

func (c *Controller) SaveToFileWithBuffer(fromFile string, toFile string, buf []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) StartSession() session.Store {
	_ = "STUB: not implemented"
	return *new(session.Store)
}

func (c *Controller) SetSession(name interface{}, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) GetSession(name interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) DelSession(name interface{}) error { _ = "STUB: not implemented"; return nil }

func (c *Controller) SessionRegenerateID() error { _ = "STUB: not implemented"; return nil }

func (c *Controller) DestroySession() error { _ = "STUB: not implemented"; return nil }

func (c *Controller) IsAjax() bool { _ = "STUB: not implemented"; return false }

func (c *Controller) GetSecureCookie(Secret, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (c *Controller) SetSecureCookie(Secret, name, value string, others ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) XSRFToken() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) CheckXSRFCookie() bool { _ = "STUB: not implemented"; return false }

func (c *Controller) XSRFFormHTML() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) GetControllerAndAction() (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}
