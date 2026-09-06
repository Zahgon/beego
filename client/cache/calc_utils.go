package cache

import (
	"github.com/beego/beego/v2/core/berror"
)

var (
	ErrIncrementOverflow = berror.Error(IncrementOverflow, "this incr invocation will overflow.")
	ErrDecrementOverflow = berror.Error(DecrementOverflow, "this decr invocation will overflow.")
	ErrNotIntegerType    = berror.Error(NotIntegerType, "item val is not (u)int (u)int32 (u)int64")
)

const (
	MinUint32 uint32 = 0
	MinUint64 uint64 = 0
)

func incr(originVal interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func decr(originVal interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }
