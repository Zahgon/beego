package bean

import (
	"context"
	"reflect"
)

const DefaultValueTagKey = "default"

type TagAutoWireBeanFactory struct {
	Adapters map[string]TypeAdapter

	FieldTagParser func(field reflect.StructField) *FieldMetadata
}

func NewTagAutoWireBeanFactory() *TagAutoWireBeanFactory { _ = "STUB: not implemented"; return nil }

func (t *TagAutoWireBeanFactory) AutoWire(ctx context.Context, appCtx ApplicationContext, bean interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TagAutoWireBeanFactory) setFloatXValue(dftValue string, bitSize int, fn string, fv reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TagAutoWireBeanFactory) setUIntXValue(dftValue string, bitSize int, fn string, fv reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TagAutoWireBeanFactory) setIntXValue(dftValue string, bitSize int, fn string, fv reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TagAutoWireBeanFactory) needInject(fValue reflect.Value) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *TagAutoWireBeanFactory) getConfig(beanValue reflect.Value) *BeanMetadata {
	_ = "STUB: not implemented"
	return nil
}
