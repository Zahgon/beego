package web

import (
	"net/http"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

type namespaceCond func(*beecontext.Context) bool

type LinkNamespace func(*Namespace)

type Namespace struct {
	prefix   string
	handlers *ControllerRegister
}

func NewNamespace(prefix string, params ...LinkNamespace) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) Cond(cond namespaceCond) *Namespace { _ = "STUB: not implemented"; return nil }

func (n *Namespace) Filter(action string, filter ...FilterFunc) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) Router(rootpath string, c ControllerInterface, mappingMethods ...string) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) AutoRouter(c ControllerInterface) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) AutoPrefix(prefix string, c ControllerInterface) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) Get(rootpath string, f HandleFunc) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) Post(rootpath string, f HandleFunc) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) Delete(rootpath string, f HandleFunc) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) Put(rootpath string, f HandleFunc) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) Head(rootpath string, f HandleFunc) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) Options(rootpath string, f HandleFunc) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) Patch(rootpath string, f HandleFunc) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) Any(rootpath string, f HandleFunc) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) Handler(rootpath string, h http.Handler) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) Include(cList ...ControllerInterface) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) CtrlGet(rootpath string, f interface{}) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) CtrlPost(rootpath string, f interface{}) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) CtrlDelete(rootpath string, f interface{}) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) CtrlPut(rootpath string, f interface{}) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) CtrlHead(rootpath string, f interface{}) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) CtrlOptions(rootpath string, f interface{}) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) CtrlPatch(rootpath string, f interface{}) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) CtrlAny(rootpath string, f interface{}) *Namespace {
	_ = "STUB: not implemented"
	return nil
}

func (n *Namespace) Namespace(ns ...*Namespace) *Namespace { _ = "STUB: not implemented"; return nil }

func AddNamespace(nl ...*Namespace) { _ = "STUB: not implemented"; return }

func addPrefix(t *Tree, prefix string) { _ = "STUB: not implemented"; return }

func NSCond(cond namespaceCond) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSBefore(filterList ...FilterFunc) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSAfter(filterList ...FilterFunc) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSInclude(cList ...ControllerInterface) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSRouter(rootpath string, c ControllerInterface, mappingMethods ...string) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSGet(rootpath string, f HandleFunc) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSPost(rootpath string, f HandleFunc) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSHead(rootpath string, f HandleFunc) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSPut(rootpath string, f HandleFunc) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSDelete(rootpath string, f HandleFunc) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSAny(rootpath string, f HandleFunc) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSOptions(rootpath string, f HandleFunc) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSPatch(rootpath string, f HandleFunc) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSCtrlGet(rootpath string, f interface{}) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSCtrlPost(rootpath string, f interface{}) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSCtrlHead(rootpath string, f interface{}) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSCtrlPut(rootpath string, f interface{}) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSCtrlDelete(rootpath string, f interface{}) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSCtrlAny(rootpath string, f interface{}) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSCtrlOptions(rootpath string, f interface{}) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSCtrlPatch(rootpath string, f interface{}) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSAutoRouter(c ControllerInterface) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSAutoPrefix(prefix string, c ControllerInterface) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSNamespace(prefix string, params ...LinkNamespace) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}

func NSHandler(rootpath string, h http.Handler) LinkNamespace {
	_ = "STUB: not implemented"
	return *new(LinkNamespace)
}
