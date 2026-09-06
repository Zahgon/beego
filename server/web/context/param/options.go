package param

type MethodParamOption func(*MethodParam)

var IsRequired MethodParamOption = func(p *MethodParam) {
	p.required = true
}

var InHeader MethodParamOption = func(p *MethodParam) {
	p.in = header
}

var InPath MethodParamOption = func(p *MethodParam) {
	p.in = path
}

var InBody MethodParamOption = func(p *MethodParam) {
	p.in = body
}

func Default(defaultValue interface{}) MethodParamOption {
	_ = "STUB: not implemented"
	return *new(MethodParamOption)
}
