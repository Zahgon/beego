package mock

import (
	"context"
	"net/http"
)

const mockCtxKey = "beego-httplib-mock"

func init() {
	InitMockSetting()
}

type Stub interface {
	Mock(cond RequestCondition, resp *http.Response, err error)
	Clear()
	MockByPath(path string, resp *http.Response, err error)
}

var mockFilter = &MockResponseFilter{}

func InitMockSetting() { _ = "STUB: not implemented"; return }

func StartMock() Stub { _ = "STUB: not implemented"; return *new(Stub) }

func CtxWithMock(ctx context.Context, mock ...*Mock) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func mockFromCtx(ctx context.Context) []*Mock { _ = "STUB: not implemented"; return nil }

type Mock struct {
	cond RequestCondition
	resp *http.Response
	err  error
}

func NewMockByPath(path string, resp *http.Response, err error) *Mock {
	_ = "STUB: not implemented"
	return nil
}

func NewMock(con RequestCondition, resp *http.Response, err error) *Mock {
	_ = "STUB: not implemented"
	return nil
}
