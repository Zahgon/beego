package config

type fakeConfigContainer struct {
	BaseConfiger
	data map[string]string
}

func (c *fakeConfigContainer) getData(key string) string { _ = "STUB: not implemented"; return "" }

func (c *fakeConfigContainer) Set(key, val string) error { _ = "STUB: not implemented"; return nil }

func (c *fakeConfigContainer) Int(key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *fakeConfigContainer) DefaultInt(key string, defaultVal int) int {
	_ = "STUB: not implemented"
	return 0
}

func (c *fakeConfigContainer) Int64(key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *fakeConfigContainer) DefaultInt64(key string, defaultVal int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *fakeConfigContainer) Bool(key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *fakeConfigContainer) DefaultBool(key string, defaultVal bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *fakeConfigContainer) Float(key string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *fakeConfigContainer) DefaultFloat(key string, defaultVal float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *fakeConfigContainer) DIY(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *fakeConfigContainer) GetSection(section string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *fakeConfigContainer) SaveConfigFile(filename string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *fakeConfigContainer) Unmarshaler(prefix string, obj interface{}, opt ...DecodeOption) error {
	_ = "STUB: not implemented"
	return nil
}

var _ Configer = new(fakeConfigContainer)

func NewFakeConfig() Configer { _ = "STUB: not implemented"; return *new(Configer) }
