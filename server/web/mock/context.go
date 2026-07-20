package mock

import (
	"net/http"
	"net/http/httptest"

	beegoCtx "github.com/beego/beego/v2/server/web/context"
)

func NewMockContext(req *http.Request) (*beegoCtx.Context, *httptest.ResponseRecorder) {
	_ = "STUB: not implemented"
	return nil, nil
}
