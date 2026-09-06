package auth

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

var defaultRealm = "Authorization Required"

func Basic(username string, password string) web.FilterFunc {
	_ = "STUB: not implemented"
	return *new(web.FilterFunc)
}

func NewBasicAuthenticator(secrets SecretProvider, Realm string) web.FilterFunc {
	_ = "STUB: not implemented"
	return *new(web.FilterFunc)
}

type SecretProvider func(user, pass string) bool

type BasicAuth struct {
	Secrets SecretProvider
	Realm   string
}

func (a *BasicAuth) CheckAuth(r *http.Request) string { _ = "STUB: not implemented"; return "" }

func (a *BasicAuth) RequireAuth(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
