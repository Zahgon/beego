package config

import (
	"fmt"
	"os"
	"sync"
)

var (
	defaultSection = "default"
	bNumComment    = []byte{'#'}
	bSemComment    = []byte{';'}
	bEmpty         = []byte{}
	bEqual         = []byte{'='}
	bDQuote        = []byte{'"'}
	sectionStart   = []byte{'['}
	sectionEnd     = []byte{']'}
	lineBreak      = "\n"
)

type IniConfig struct{}

func (ini *IniConfig) Parse(name string) (Configer, error) {
	_ = "STUB: not implemented"
	return *new(Configer), nil
}

func (ini *IniConfig) parseFile(name string) (*IniConfigContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ini *IniConfig) parseData(dir string, data []byte) (*IniConfigContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ini *IniConfig) ParseData(data []byte) (Configer, error) {
	_ = "STUB: not implemented"
	return *new(Configer), nil
}

type IniConfigContainer struct {
	BaseConfiger
	data           map[string]map[string]string
	sectionComment map[string]string
	keyComment     map[string]string
	sync.RWMutex
}

func (c *IniConfigContainer) Bool(key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *IniConfigContainer) DefaultBool(key string, defaultVal bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *IniConfigContainer) Int(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *IniConfigContainer) DefaultInt(key string, defaultVal int) int {
	_ = "STUB: not implemented"
	return 0
}

func (c *IniConfigContainer) Int64(key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *IniConfigContainer) DefaultInt64(key string, defaultVal int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *IniConfigContainer) Float(key string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *IniConfigContainer) DefaultFloat(key string, defaultVal float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *IniConfigContainer) String(key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *IniConfigContainer) DefaultString(key string, defaultVal string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *IniConfigContainer) Strings(key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *IniConfigContainer) DefaultStrings(key string, defaultVal []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *IniConfigContainer) GetSection(section string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *IniConfigContainer) SaveConfigFile(filename string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *IniConfigContainer) Set(key, val string) error { _ = "STUB: not implemented"; return nil }

func (c *IniConfigContainer) DIY(key string) (v interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *IniConfigContainer) getdata(key string) string { _ = "STUB: not implemented"; return "" }

func (c *IniConfigContainer) Unmarshaler(prefix string, obj interface{}, opt ...DecodeOption) error {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	Register("ini", &IniConfig{})

	err := InitGlobalInstance("ini", "conf/app.conf")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "init global config instance failed. If you do not use this, just ignore it. ", err)
	}
}
