package config

var globalInstance Configer

func InitGlobalInstance(name string, cfg string) error { _ = "STUB: not implemented"; return nil }

func Set(key, val string) error { _ = "STUB: not implemented"; return nil }

func String(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func Strings(key string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Int(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func Int64(key string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func Bool(key string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func Float(key string) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func DefaultString(key string, defaultVal string) string { _ = "STUB: not implemented"; return "" }

func DefaultStrings(key string, defaultVal []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func DefaultInt(key string, defaultVal int) int { _ = "STUB: not implemented"; return 0 }

func DefaultInt64(key string, defaultVal int64) int64 { _ = "STUB: not implemented"; return 0 }

func DefaultBool(key string, defaultVal bool) bool { _ = "STUB: not implemented"; return false }

func DefaultFloat(key string, defaultVal float64) float64 { _ = "STUB: not implemented"; return 0 }

func DIY(key string) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func GetSection(section string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Unmarshaler(prefix string, obj interface{}, opt ...DecodeOption) error {
	_ = "STUB: not implemented"
	return nil
}

func Sub(key string) (Configer, error) { _ = "STUB: not implemented"; return *new(Configer), nil }

func OnChange(key string, fn func(value string)) { _ = "STUB: not implemented"; return }

func SaveConfigFile(filename string) error { _ = "STUB: not implemented"; return nil }
