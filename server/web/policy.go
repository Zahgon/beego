package web

import (
	"github.com/beego/beego/v2/server/web/context"
)

type PolicyFunc func(*context.Context)

func (p *ControllerRegister) FindPolicy(cont *context.Context) []PolicyFunc {
	_ = "STUB: not implemented"
	return nil
}

func (p *ControllerRegister) addToPolicy(method, pattern string, r ...PolicyFunc) {
	_ = "STUB: not implemented"
	return
}

func Policy(pattern, method string, policy ...PolicyFunc) { _ = "STUB: not implemented"; return }

func (p *ControllerRegister) execPolicy(cont *context.Context, urlPath string) (started bool) {
	_ = "STUB: not implemented"
	return false
}
