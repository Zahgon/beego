package toml

import (
	"github.com/pelletier/go-toml"

	"github.com/beego/beego/v2/core/config"
)

const keySeparator = "."

type Config struct {
	tree *toml.Tree
}

func (c *Config) Parse(filename string) (config.Configer, error) {
	_ = "STUB: not implemented"
	return *new(config.Configer), nil
}

func (c *Config) ParseData(data []byte) (config.Configer, error) {
	_ = "STUB: not implemented"
	return *new(config.Configer), nil
}

type configContainer struct {
	t *toml.Tree
}

func (c *configContainer) Set(key, val string) error { _ = "STUB: not implemented"; return nil }

func (c *configContainer) String(key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *configContainer) Strings(key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *configContainer) Int(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *configContainer) Int64(key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *configContainer) Bool(key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *configContainer) Float(key string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *configContainer) DefaultString(key string, defaultVal string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *configContainer) DefaultStrings(key string, defaultVal []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *configContainer) DefaultInt(key string, defaultVal int) int {
	_ = "STUB: not implemented"
	return 0
}

func (c *configContainer) DefaultInt64(key string, defaultVal int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *configContainer) DefaultBool(key string, defaultVal bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *configContainer) DefaultFloat(key string, defaultVal float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *configContainer) DIY(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *configContainer) GetSection(section string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *configContainer) Unmarshaler(prefix string, obj interface{}, opt ...config.DecodeOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *configContainer) Sub(key string) (config.Configer, error) {
	_ = "STUB: not implemented"
	return *new(config.Configer), nil
}

func (c *configContainer) OnChange(key string, fn func(value string)) {
	_ = "STUB: not implemented"
	return
}

func (c *configContainer) SaveConfigFile(filename string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *configContainer) get(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func subTree(t *toml.Tree, path []string) (*toml.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func init() {
	config.Register("toml", &Config{})
}
