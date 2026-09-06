package session

import "net/http"

type ManagerConfig struct {
	EnableSetCookie         bool          `json:"enableSetCookie,omitempty"`
	DisableHTTPOnly         bool          `json:"disableHTTPOnly"`
	Secure                  bool          `json:"secure"`
	EnableSidInHTTPHeader   bool          `json:"EnableSidInHTTPHeader"`
	EnableSidInURLQuery     bool          `json:"EnableSidInURLQuery"`
	CookieName              string        `json:"cookieName"`
	Gclifetime              int64         `json:"gclifetime"`
	Maxlifetime             int64         `json:"maxLifetime"`
	CookieLifeTime          int           `json:"cookieLifeTime"`
	ProviderConfig          string        `json:"providerConfig"`
	Domain                  string        `json:"domain"`
	SessionIDLength         int64         `json:"sessionIDLength"`
	SessionNameInHTTPHeader string        `json:"SessionNameInHTTPHeader"`
	SessionIDPrefix         string        `json:"sessionIDPrefix"`
	CookieSameSite          http.SameSite `json:"cookieSameSite"`
}

func (c *ManagerConfig) Opts(opts ...ManagerConfigOpt) { _ = "STUB: not implemented"; return }

type ManagerConfigOpt func(config *ManagerConfig)

func NewManagerConfig(opts ...ManagerConfigOpt) *ManagerConfig {
	_ = "STUB: not implemented"
	return nil
}

func CfgCookieName(cookieName string) ManagerConfigOpt {
	_ = "STUB: not implemented"
	return *new(ManagerConfigOpt)
}

func CfgSessionIdLength(length int64) ManagerConfigOpt {
	_ = "STUB: not implemented"
	return *new(ManagerConfigOpt)
}

func CfgSessionIdPrefix(prefix string) ManagerConfigOpt {
	_ = "STUB: not implemented"
	return *new(ManagerConfigOpt)
}

func CfgSetCookie(enable bool) ManagerConfigOpt {
	_ = "STUB: not implemented"
	return *new(ManagerConfigOpt)
}

func CfgGcLifeTime(lifeTime int64) ManagerConfigOpt {
	_ = "STUB: not implemented"
	return *new(ManagerConfigOpt)
}

func CfgMaxLifeTime(lifeTime int64) ManagerConfigOpt {
	_ = "STUB: not implemented"
	return *new(ManagerConfigOpt)
}

func CfgCookieLifeTime(lifeTime int) ManagerConfigOpt {
	_ = "STUB: not implemented"
	return *new(ManagerConfigOpt)
}

func CfgProviderConfig(providerConfig string) ManagerConfigOpt {
	_ = "STUB: not implemented"
	return *new(ManagerConfigOpt)
}

func CfgDomain(domain string) ManagerConfigOpt {
	_ = "STUB: not implemented"
	return *new(ManagerConfigOpt)
}

func CfgSessionIdInHTTPHeader(enable bool) ManagerConfigOpt {
	_ = "STUB: not implemented"
	return *new(ManagerConfigOpt)
}

func CfgSetSessionNameInHTTPHeader(name string) ManagerConfigOpt {
	_ = "STUB: not implemented"
	return *new(ManagerConfigOpt)
}

func CfgEnableSidInURLQuery(enable bool) ManagerConfigOpt {
	_ = "STUB: not implemented"
	return *new(ManagerConfigOpt)
}

func CfgHTTPOnly(HTTPOnly bool) ManagerConfigOpt {
	_ = "STUB: not implemented"
	return *new(ManagerConfigOpt)
}

func CfgSecure(Enable bool) ManagerConfigOpt {
	_ = "STUB: not implemented"
	return *new(ManagerConfigOpt)
}

func CfgSameSite(sameSite http.SameSite) ManagerConfigOpt {
	_ = "STUB: not implemented"
	return *new(ManagerConfigOpt)
}
