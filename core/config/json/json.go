package json

import (
	"sync"

	"github.com/beego/beego/v2/core/config"
)

type JSONConfig struct{}

func (js *JSONConfig) Parse(filename string) (config.Configer, error) {
	_ = "STUB: not implemented"
	return *new(config.Configer), nil
}

func (js *JSONConfig) ParseData(data []byte) (config.Configer, error) {
	_ = "STUB: not implemented"
	return *new(config.Configer), nil
}

type JSONConfigContainer struct {
	data map[string]interface{}
	sync.RWMutex
}

func (c *JSONConfigContainer) Unmarshaler(prefix string, obj interface{}, opt ...config.DecodeOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *JSONConfigContainer) Sub(key string) (config.Configer, error) {
	_ = "STUB: not implemented"
	return *new(config.Configer), nil
}

func (c *JSONConfigContainer) sub(key string) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *JSONConfigContainer) OnChange(key string, fn func(value string)) {
	_ = "STUB: not implemented"
	return
}

func (c *JSONConfigContainer) Bool(key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *JSONConfigContainer) DefaultBool(key string, defaultVal bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *JSONConfigContainer) Int(key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *JSONConfigContainer) DefaultInt(key string, defaultVal int) int {
	_ = "STUB: not implemented"
	return 0
}

func (c *JSONConfigContainer) Int64(key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *JSONConfigContainer) DefaultInt64(key string, defaultVal int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *JSONConfigContainer) Float(key string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *JSONConfigContainer) DefaultFloat(key string, defaultVal float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *JSONConfigContainer) String(key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *JSONConfigContainer) DefaultString(key string, defaultVal string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *JSONConfigContainer) Strings(key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *JSONConfigContainer) DefaultStrings(key string, defaultVal []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *JSONConfigContainer) GetSection(section string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *JSONConfigContainer) SaveConfigFile(filename string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *JSONConfigContainer) Set(key, val string) error { _ = "STUB: not implemented"; return nil }

func (c *JSONConfigContainer) DIY(key string) (v interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *JSONConfigContainer) getData(key string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	config.Register("json", &JSONConfig{})
}
