package cors

import (
	"time"

	"github.com/beego/beego/v2/server/web"
)

const (
	headerAllowOrigin      = "Access-Control-Allow-Origin"
	headerAllowCredentials = "Access-Control-Allow-Credentials"
	headerAllowHeaders     = "Access-Control-Allow-Headers"
	headerAllowMethods     = "Access-Control-Allow-Methods"
	headerExposeHeaders    = "Access-Control-Expose-Headers"
	headerMaxAge           = "Access-Control-Max-Age"

	headerOrigin         = "Origin"
	headerRequestMethod  = "Access-Control-Request-Method"
	headerRequestHeaders = "Access-Control-Request-Headers"
)

var (
	defaultAllowHeaders = []string{"Origin", "Accept", "Content-Type", "Authorization"}

	allowOriginPatterns = []string{}
)

type Options struct {
	AllowAllOrigins bool

	AllowCredentials bool

	AllowOrigins []string

	AllowMethods []string

	AllowHeaders []string

	ExposeHeaders []string

	MaxAge time.Duration
}

func (o *Options) Header(origin string) (headers map[string]string) {
	_ = "STUB: not implemented"
	return nil
}

func (o *Options) PreflightHeader(origin, rMethod, rHeaders string) (headers map[string]string) {
	_ = "STUB: not implemented"
	return nil
}

func (o *Options) IsOriginAllowed(origin string) (allowed bool) {
	_ = "STUB: not implemented"
	return false
}

func Allow(opts *Options) web.FilterFunc { _ = "STUB: not implemented"; return *new(web.FilterFunc) }
