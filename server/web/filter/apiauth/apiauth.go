package apiauth

import (
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type AppIDToAppSecret func(string) string

func APIBasicAuth(appid, appkey string) web.FilterFunc {
	_ = "STUB: not implemented"
	return *new(web.FilterFunc)
}

func APISecretAuth(f AppIDToAppSecret, timeout int) web.FilterFunc {
	_ = "STUB: not implemented"
	return *new(web.FilterFunc)
}

func Signature(appsecret, method string, params url.Values, RequestURL string) (result string) {
	_ = "STUB: not implemented"
	return ""
}
