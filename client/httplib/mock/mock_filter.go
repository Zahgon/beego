package mock

import (
	"net/http"

	"github.com/beego/beego/v2/client/httplib"
)

type MockResponseFilter struct {
	ms []*Mock
}

func NewMockResponseFilter() *MockResponseFilter { _ = "STUB: not implemented"; return nil }

func (m *MockResponseFilter) FilterChain(next httplib.Filter) httplib.Filter {
	_ = "STUB: not implemented"
	return *new(httplib.Filter)
}

func (m *MockResponseFilter) MockByPath(path string, resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return
}

func (m *MockResponseFilter) Clear() { _ = "STUB: not implemented"; return }

func (m *MockResponseFilter) Mock(cond RequestCondition, resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return
}
