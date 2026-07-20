package web

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/utils"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/session"
)

type Config struct {
	AppName string

	RunMode string

	RouterCaseSensitive bool

	RecoverPanic bool

	CopyRequestBody bool

	EnableGzip bool

	EnableErrorsShow bool

	EnableErrorsRender bool

	ServerName string

	RecoverFunc func(*context.Context, *Config)

	MaxMemory int64

	MaxUploadSize int64

	Listen Listen

	WebConfig WebConfig

	Log LogConfig
}

type Listen struct {
	Graceful bool

	ListenTCP4 bool

	EnableHTTP bool

	AutoTLS bool

	EnableHTTPS bool

	EnableMutualHTTPS bool

	EnableAdmin bool

	EnableFcgi bool

	EnableStdIo bool

	ServerTimeOut int64

	HTTPAddr string

	HTTPPort int

	Domains []string

	TLSCacheDir string

	HTTPSAddr string

	HTTPSPort int

	HTTPSCertFile string

	HTTPSKeyFile string

	TrustCaFile string

	AdminAddr string

	AdminPort int

	ClientAuth int
}

type WebConfig struct {
	AutoRender bool

	EnableDocs bool

	EnableXSRF bool

	DirectoryIndex bool

	FlashName string

	FlashSeparator string

	StaticDir map[string]string

	StaticExtensionsToGzip []string

	StaticCacheFileSize int

	StaticCacheFileNum int

	TemplateLeft string

	TemplateRight string

	ViewsPath string

	CommentRouterPath string

	XSRFKey string

	XSRFExpire int

	Session SessionConfig
}

type SessionConfig struct {
	SessionOn bool

	SessionAutoSetCookie bool

	SessionDisableHTTPOnly bool

	SessionEnableSidInHTTPHeader bool

	SessionEnableSidInURLQuery bool

	SessionProvider string

	SessionName string

	SessionGCMaxLifetime int64

	SessionProviderConfig string

	SessionCookieLifeTime int

	SessionDomain string

	SessionNameInHTTPHeader string

	SessionCookieSameSite http.SameSite

	SessionIDPrefix string
}

type LogConfig struct {
	AccessLogs bool

	EnableStaticLogs bool

	FileLineNum bool

	AccessLogsFormat string

	Outputs map[string]string
}

var (
	BConfig *Config

	AppConfig *beegoAppConfig

	AppPath string

	GlobalSessions *session.Manager

	appConfigPath string

	appConfigProvider = "ini"

	WorkPath string
)

func init() {
	BConfig = newBConfig()
	var err error
	if AppPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}
	WorkPath, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	filename := "app.conf"
	if os.Getenv("BEEGO_RUNMODE") != "" {
		filename = os.Getenv("BEEGO_RUNMODE") + ".app.conf"
	}
	appConfigPath = filepath.Join(WorkPath, "conf", filename)
	if !utils.FileExists(appConfigPath) {
		appConfigPath = filepath.Join(AppPath, "conf", filename)
		if !utils.FileExists(appConfigPath) {
			AppConfig = &beegoAppConfig{innerConfig: config.NewFakeConfig()}
			return
		}
	}
	if err = parseConfig(appConfigPath); err != nil {
		panic(err)
	}
}

func defaultRecoverPanic(ctx *context.Context, cfg *Config) { _ = "STUB: not implemented"; return }

func newBConfig() *Config { _ = "STUB: not implemented"; return nil }

func parseConfig(appConfigPath string) (err error) { _ = "STUB: not implemented"; return nil }

func assignConfig(ac config.Configer) error { _ = "STUB: not implemented"; return nil }

func parseConfigForV1(ac config.Configer) { _ = "STUB: not implemented"; return }

func assignSingleConfig(p interface{}, ac config.Configer) { _ = "STUB: not implemented"; return }

func LoadAppConfig(adapterName, configPath string) error { _ = "STUB: not implemented"; return nil }

type beegoAppConfig struct {
	config.BaseConfiger
	innerConfig config.Configer
}

func newAppConfig(appConfigProvider, appConfigPath string) (*beegoAppConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *beegoAppConfig) Unmarshaler(prefix string, obj interface{}, opt ...config.DecodeOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *beegoAppConfig) Set(key, val string) error { _ = "STUB: not implemented"; return nil }

func (b *beegoAppConfig) String(key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (b *beegoAppConfig) Strings(key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *beegoAppConfig) Int(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *beegoAppConfig) Int64(key string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *beegoAppConfig) Bool(key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (b *beegoAppConfig) Float(key string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *beegoAppConfig) DefaultString(key string, defaultVal string) string {
	_ = "STUB: not implemented"
	return ""
}

func (b *beegoAppConfig) DefaultStrings(key string, defaultVal []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (b *beegoAppConfig) DefaultInt(key string, defaultVal int) int {
	_ = "STUB: not implemented"
	return 0
}

func (b *beegoAppConfig) DefaultInt64(key string, defaultVal int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (b *beegoAppConfig) DefaultBool(key string, defaultVal bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (b *beegoAppConfig) DefaultFloat(key string, defaultVal float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (b *beegoAppConfig) DIY(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *beegoAppConfig) GetSection(section string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *beegoAppConfig) SaveConfigFile(filename string) error {
	_ = "STUB: not implemented"
	return nil
}
