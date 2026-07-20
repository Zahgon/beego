package session

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/session"
)

func Session(providerType session.ProviderType, options ...session.ManagerConfigOpt) web.FilterChain {
	_ = "STUB: not implemented"
	return *new(web.FilterChain)
}
