package mock

import (
	"github.com/beego/beego/v2/client/orm"
)

var stub = newOrmStub()

func init() {
	orm.AddGlobalFilterChain(stub.FilterChain)
}

type Stub interface {
	Mock(m *Mock)
	Clear()
}

type OrmStub struct {
	ms []*Mock
}

func StartMock() Stub { _ = "STUB: not implemented"; return *new(Stub) }

func newOrmStub() *OrmStub { _ = "STUB: not implemented"; return nil }

func (o *OrmStub) Mock(m *Mock) { _ = "STUB: not implemented"; return }

func (o *OrmStub) Clear() { _ = "STUB: not implemented"; return }

func (o *OrmStub) FilterChain(next orm.Filter) orm.Filter {
	_ = "STUB: not implemented"
	return *new(orm.Filter)
}
