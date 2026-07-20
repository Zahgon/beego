package utils

func SelfPath() string { _ = "STUB: not implemented"; return "" }

func SelfDir() string { _ = "STUB: not implemented"; return "" }

func FileExists(name string) bool { _ = "STUB: not implemented"; return false }

func SearchFile(filename string, paths ...string) (fullpath string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GrepFile(patten string, filename string) (lines []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
