package utils

type KV interface {
	GetKey() interface{}
	GetValue() interface{}
}

type SimpleKV struct {
	Key   interface{}
	Value interface{}
}

var _ KV = new(SimpleKV)

func (s *SimpleKV) GetKey() interface{} { _ = "STUB: not implemented"; return nil }

func (s *SimpleKV) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

type KVs interface {
	GetValueOr(key interface{}, defValue interface{}) interface{}
	Contains(key interface{}) bool
	IfContains(key interface{}, action func(value interface{})) KVs
}

type SimpleKVs struct {
	kvs map[interface{}]interface{}
}

var _ KVs = new(SimpleKVs)

func (kvs *SimpleKVs) GetValueOr(key interface{}, defValue interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (kvs *SimpleKVs) Contains(key interface{}) bool { _ = "STUB: not implemented"; return false }

func (kvs *SimpleKVs) IfContains(key interface{}, action func(value interface{})) KVs {
	_ = "STUB: not implemented"
	return *new(KVs)
}

func NewKVs(kvs ...KV) KVs { _ = "STUB: not implemented"; return *new(KVs) }
