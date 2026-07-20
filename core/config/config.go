package config

import (
	"context"
)

type Configer interface {
	Set(key, val string) error

	String(key string) (string, error)

	Strings(key string) ([]string, error)
	Int(key string) (int, error)
	Int64(key string) (int64, error)
	Bool(key string) (bool, error)
	Float(key string) (float64, error)

	DefaultString(key string, defaultVal string) string

	DefaultStrings(key string, defaultVal []string) []string
	DefaultInt(key string, defaultVal int) int
	DefaultInt64(key string, defaultVal int64) int64
	DefaultBool(key string, defaultVal bool) bool
	DefaultFloat(key string, defaultVal float64) float64

	DIY(key string) (interface{}, error)

	GetSection(section string) (map[string]string, error)

	Unmarshaler(prefix string, obj interface{}, opt ...DecodeOption) error
	Sub(key string) (Configer, error)
	OnChange(key string, fn func(value string))
	SaveConfigFile(filename string) error
}

type BaseConfiger struct {
	reader func(ctx context.Context, key string) (string, error)
}

func NewBaseConfiger(reader func(ctx context.Context, key string) (string, error)) BaseConfiger {
	_ = "STUB: not implemented"
	return *new(BaseConfiger)
}

func (c *BaseConfiger) Int(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *BaseConfiger) Int64(key string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *BaseConfiger) Bool(key string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (c *BaseConfiger) Float(key string) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *BaseConfiger) DefaultString(key string, defaultVal string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *BaseConfiger) DefaultStrings(key string, defaultVal []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *BaseConfiger) DefaultInt(key string, defaultVal int) int {
	_ = "STUB: not implemented"
	return 0
}

func (c *BaseConfiger) DefaultInt64(key string, defaultVal int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *BaseConfiger) DefaultBool(key string, defaultVal bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *BaseConfiger) DefaultFloat(key string, defaultVal float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *BaseConfiger) String(key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *BaseConfiger) Strings(key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*BaseConfiger) Sub(string) (Configer, error) {
	_ = "STUB: not implemented"
	return *new(Configer), nil
}

func (*BaseConfiger) OnChange(_ string, _ func(value string)) { _ = "STUB: not implemented"; return }

type Config interface {
	Parse(key string) (Configer, error)
	ParseData(data []byte) (Configer, error)
}

var adapters = make(map[string]Config)

func Register(name string, adapter Config) { _ = "STUB: not implemented"; return }

func NewConfig(adapterName, filename string) (Configer, error) {
	_ = "STUB: not implemented"
	return *new(Configer), nil
}

func NewConfigData(adapterName string, data []byte) (Configer, error) {
	_ = "STUB: not implemented"
	return *new(Configer), nil
}

func ExpandValueEnvForMap(m map[string]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func ExpandValueEnv(value string) (realValue string) { _ = "STUB: not implemented"; return "" }

func ParseBool(val interface{}) (value bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func ToString(x interface{}) string { _ = "STUB: not implemented"; return "" }

type DecodeOption func(options decodeOptions)

type decodeOptions struct{}
