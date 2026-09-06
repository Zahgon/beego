package yaml

import (
	"sync"

	"github.com/beego/beego/v2/core/config"
)

type Config struct{}

func (*Config) Parse(filename string) (y config.Configer, err error) {
	_ = "STUB: not implemented"
	return *new(config.Configer), nil
}

func (*Config) ParseData(data []byte) (config.Configer, error) {
	_ = "STUB: not implemented"
	return *new(config.Configer), nil
}

func ReadYmlReader(path string) (cnf map[string]interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseYML(buf []byte) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ConfigContainer struct {
	data map[string]interface{}
	sync.RWMutex
}

func (c *ConfigContainer) Unmarshaler(prefix string, obj interface{}, _ ...config.DecodeOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigContainer) Sub(key string) (config.Configer, error) {
	_ = "STUB: not implemented"
	return *new(config.Configer), nil
}

func (c *ConfigContainer) subMap(key string) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*ConfigContainer) OnChange(_ string, _ func(value string)) { _ = "STUB: not implemented"; return }

func (c *ConfigContainer) Bool(key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *ConfigContainer) DefaultBool(key string, defaultVal bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ConfigContainer) Int(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *ConfigContainer) DefaultInt(key string, defaultVal int) int {
	_ = "STUB: not implemented"
	return 0
}

func (c *ConfigContainer) Int64(key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *ConfigContainer) DefaultInt64(key string, defaultVal int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *ConfigContainer) Float(key string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *ConfigContainer) DefaultFloat(key string, defaultVal float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *ConfigContainer) String(key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *ConfigContainer) DefaultString(key string, defaultVal string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *ConfigContainer) Strings(key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConfigContainer) DefaultStrings(key string, defaultVal []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigContainer) GetSection(section string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConfigContainer) SaveConfigFile(filename string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigContainer) Set(key, val string) error { _ = "STUB: not implemented"; return nil }

func (c *ConfigContainer) DIY(key string) (v interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConfigContainer) getData(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func init() {
	config.Register("yaml", &Config{})
}
