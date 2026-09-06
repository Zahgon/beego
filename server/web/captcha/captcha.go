package captcha

import (
	context2 "context"
	"html/template"
	"net/http"
	"time"

	"github.com/beego/beego/v2/server/web/context"
)

var defaultChars = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

const (
	challengeNums    = 6
	expiration       = 600 * time.Second
	fieldIDName      = "captcha_id"
	fieldCaptchaName = "captcha"
	cachePrefix      = "captcha_"
	defaultURLPrefix = "/captcha/"
)

type Captcha struct {
	store Storage

	URLPrefix string

	FieldIDName string

	FieldCaptchaName string

	StdWidth  int
	StdHeight int

	ChallengeNums int

	Expiration time.Duration

	CachePrefix string
}

func (c *Captcha) key(id string) string { _ = "STUB: not implemented"; return "" }

func (c *Captcha) genRandChars() []byte { _ = "STUB: not implemented"; return nil }

func (c *Captcha) Handler(ctx *context.Context) { _ = "STUB: not implemented"; return }

func (c *Captcha) CreateCaptchaHTML() template.HTML {
	_ = "STUB: not implemented"
	return *new(template.HTML)
}

func (c *Captcha) CreateCaptcha() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *Captcha) VerifyReq(req *http.Request) bool { _ = "STUB: not implemented"; return false }

func (c *Captcha) Verify(id string, challenge string) (success bool) {
	_ = "STUB: not implemented"
	return false
}

func NewCaptcha(urlPrefix string, store Storage) *Captcha { _ = "STUB: not implemented"; return nil }

func NewWithFilter(urlPrefix string, store Storage) *Captcha { _ = "STUB: not implemented"; return nil }

type Storage interface {
	Get(ctx context2.Context, key string) (interface{}, error)

	Put(ctx context2.Context, key string, val interface{}, timeout time.Duration) error

	Delete(ctx context2.Context, key string) error
}
