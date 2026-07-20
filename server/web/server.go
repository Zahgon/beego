package web

import (
	"net/http"
	"time"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

var BeeApp *HttpServer

func init() {

	BeeApp = NewHttpSever()
}

type HttpServer struct {
	Handlers           *ControllerRegister
	Server             *http.Server
	Cfg                *Config
	LifeCycleCallbacks []LifeCycleCallback
}

func NewHttpSever() *HttpServer { _ = "STUB: not implemented"; return nil }

func NewHttpServerWithCfg(cfg *Config) *HttpServer { _ = "STUB: not implemented"; return nil }

type MiddleWare func(http.Handler) http.Handler

type LifeCycleCallback interface {
	AfterStart(app *HttpServer)
	BeforeShutdown(app *HttpServer)
}

func (app *HttpServer) Run(addr string, mws ...MiddleWare) { _ = "STUB: not implemented"; return }

func Router(rootpath string, c ControllerInterface, mappingMethods ...string) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func RouterWithOpts(rootpath string, c ControllerInterface, opts ...ControllerOption) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func (app *HttpServer) Router(rootPath string, c ControllerInterface, mappingMethods ...string) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func (app *HttpServer) RouterWithOpts(rootPath string, c ControllerInterface, opts ...ControllerOption) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func UnregisterFixedRoute(fixedRoute string, method string) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func (app *HttpServer) UnregisterFixedRoute(fixedRoute string, method string) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func findAndRemoveTree(paths []string, entryPointTree *Tree, method string) {
	_ = "STUB: not implemented"
	return
}

func findAndRemoveSingleTree(entryPointTree *Tree) { _ = "STUB: not implemented"; return }

func Include(cList ...ControllerInterface) *HttpServer { _ = "STUB: not implemented"; return nil }

func (app *HttpServer) Include(cList ...ControllerInterface) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func RESTRouter(rootpath string, c ControllerInterface) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func (app *HttpServer) RESTRouter(rootpath string, c ControllerInterface) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func AutoRouter(c ControllerInterface) *HttpServer { _ = "STUB: not implemented"; return nil }

func (app *HttpServer) AutoRouter(c ControllerInterface) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func AutoPrefix(prefix string, c ControllerInterface) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func (app *HttpServer) AutoPrefix(prefix string, c ControllerInterface) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func CtrlGet(rootpath string, f interface{}) { _ = "STUB: not implemented"; return }

func (app *HttpServer) CtrlGet(rootpath string, f interface{}) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func CtrlPost(rootpath string, f interface{}) { _ = "STUB: not implemented"; return }

func (app *HttpServer) CtrlPost(rootpath string, f interface{}) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func CtrlHead(rootpath string, f interface{}) { _ = "STUB: not implemented"; return }

func (app *HttpServer) CtrlHead(rootpath string, f interface{}) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func CtrlPut(rootpath string, f interface{}) { _ = "STUB: not implemented"; return }

func (app *HttpServer) CtrlPut(rootpath string, f interface{}) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func CtrlPatch(rootpath string, f interface{}) { _ = "STUB: not implemented"; return }

func (app *HttpServer) CtrlPatch(rootpath string, f interface{}) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func CtrlDelete(rootpath string, f interface{}) { _ = "STUB: not implemented"; return }

func (app *HttpServer) CtrlDelete(rootpath string, f interface{}) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func CtrlOptions(rootpath string, f interface{}) { _ = "STUB: not implemented"; return }

func (app *HttpServer) CtrlOptions(rootpath string, f interface{}) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func CtrlAny(rootpath string, f interface{}) { _ = "STUB: not implemented"; return }

func (app *HttpServer) CtrlAny(rootpath string, f interface{}) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func Get(rootpath string, f HandleFunc) *HttpServer { _ = "STUB: not implemented"; return nil }

func (app *HttpServer) Get(rootpath string, f HandleFunc) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func Post(rootpath string, f HandleFunc) *HttpServer { _ = "STUB: not implemented"; return nil }

func (app *HttpServer) Post(rootpath string, f HandleFunc) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func Delete(rootpath string, f HandleFunc) *HttpServer { _ = "STUB: not implemented"; return nil }

func (app *HttpServer) Delete(rootpath string, f HandleFunc) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func Put(rootpath string, f HandleFunc) *HttpServer { _ = "STUB: not implemented"; return nil }

func (app *HttpServer) Put(rootpath string, f HandleFunc) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func Head(rootpath string, f HandleFunc) *HttpServer { _ = "STUB: not implemented"; return nil }

func (app *HttpServer) Head(rootpath string, f HandleFunc) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func Options(rootpath string, f HandleFunc) *HttpServer { _ = "STUB: not implemented"; return nil }

func (app *HttpServer) Options(rootpath string, f HandleFunc) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func Patch(rootpath string, f HandleFunc) *HttpServer { _ = "STUB: not implemented"; return nil }

func (app *HttpServer) Patch(rootpath string, f HandleFunc) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func Any(rootpath string, f HandleFunc) *HttpServer { _ = "STUB: not implemented"; return nil }

func (app *HttpServer) Any(rootpath string, f HandleFunc) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func Handler(rootpath string, h http.Handler, options ...interface{}) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func (app *HttpServer) Handler(rootpath string, h http.Handler, options ...interface{}) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func InsertFilter(pattern string, pos int, filter FilterFunc, opts ...FilterOpt) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func (app *HttpServer) InsertFilter(pattern string, pos int, filter FilterFunc, opts ...FilterOpt) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func InsertFilterChain(pattern string, filterChain FilterChain, opts ...FilterOpt) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func (app *HttpServer) InsertFilterChain(pattern string, filterChain FilterChain, opts ...FilterOpt) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}

func (app *HttpServer) initAddr(addr string) { _ = "STUB: not implemented"; return }

func (app *HttpServer) LogAccess(ctx *beecontext.Context, startTime *time.Time, statusCode int) {
	_ = "STUB: not implemented"
	return
}

func (app *HttpServer) PrintTree() M { _ = "STUB: not implemented"; return *new(M) }

func printTree(resultList *[][]string, t *Tree) { _ = "STUB: not implemented"; return }

func (app *HttpServer) reportFilter() M { _ = "STUB: not implemented"; return *new(M) }
