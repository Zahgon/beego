package param

type MethodParam struct {
	name         string
	in           paramType
	required     bool
	defaultValue string
}

type paramType byte

const (
	param paramType = iota
	path
	body
	header
)

func New(name string, opts ...MethodParamOption) *MethodParam {
	_ = "STUB: not implemented"
	return nil
}

func newParam(name string, parser paramParser, opts []MethodParamOption) (param *MethodParam) {
	_ = "STUB: not implemented"
	return nil
}

func Make(list ...*MethodParam) []*MethodParam { _ = "STUB: not implemented"; return nil }

func (mp *MethodParam) String() string { _ = "STUB: not implemented"; return "" }
