package hints

import (
	"github.com/beego/beego/v2/core/utils"
)

const (
	KeyForceIndex = iota
	KeyUseIndex
	KeyIgnoreIndex
	KeyForUpdate
	KeyLimit
	KeyOffset
	KeyOrderBy
	KeyRelDepth
)

type Hint struct {
	key   interface{}
	value interface{}
}

var _ utils.KV = new(Hint)

func (s *Hint) GetKey() interface{} { _ = "STUB: not implemented"; return nil }

func (s *Hint) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

var _ utils.KV = new(Hint)

func ForceIndex(indexes ...string) *Hint { _ = "STUB: not implemented"; return nil }

func UseIndex(indexes ...string) *Hint { _ = "STUB: not implemented"; return nil }

func IgnoreIndex(indexes ...string) *Hint { _ = "STUB: not implemented"; return nil }

func ForUpdate() *Hint { _ = "STUB: not implemented"; return nil }

func DefaultRelDepth() *Hint { _ = "STUB: not implemented"; return nil }

func RelDepth(d int) *Hint { _ = "STUB: not implemented"; return nil }

func Limit(d int64) *Hint { _ = "STUB: not implemented"; return nil }

func Offset(d int64) *Hint { _ = "STUB: not implemented"; return nil }

func OrderBy(s string) *Hint { _ = "STUB: not implemented"; return nil }

func NewHint(key interface{}, value interface{}) *Hint { _ = "STUB: not implemented"; return nil }
