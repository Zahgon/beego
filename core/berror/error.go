package berror

const errFmt = "ERROR-%d, %s"

func Error(c Code, msg string) error { _ = "STUB: not implemented"; return nil }

func Errorf(c Code, format string, a ...interface{}) error { _ = "STUB: not implemented"; return nil }

func Wrap(err error, c Code, msg string) error { _ = "STUB: not implemented"; return nil }

func Wrapf(err error, c Code, format string, a ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func FromError(err error) (Code, bool) { _ = "STUB: not implemented"; return *new(Code), false }
