package authz

import (
	"net/http"

	"github.com/casbin/casbin"

	"github.com/beego/beego/v2/server/web"
)

func NewAuthorizer(e *casbin.Enforcer) web.FilterFunc {
	_ = "STUB: not implemented"
	return *new(web.FilterFunc)
}

type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

func (a *BasicAuthorizer) GetUserName(r *http.Request) string { _ = "STUB: not implemented"; return "" }

func (a *BasicAuthorizer) CheckPermission(r *http.Request) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *BasicAuthorizer) RequirePermission(w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}
