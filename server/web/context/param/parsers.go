package param

import (
	"reflect"
)

type paramParser interface {
	parse(value string, toType reflect.Type) (interface{}, error)
}

func getParser(param *MethodParam, t reflect.Type) paramParser {
	_ = "STUB: not implemented"
	return *new(paramParser)
}

type parserFunc func(value string, toType reflect.Type) (interface{}, error)

func (f parserFunc) parse(value string, toType reflect.Type) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type boolParser struct{}

func (p boolParser) parse(value string, toType reflect.Type) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type stringParser struct{}

func (p stringParser) parse(value string, toType reflect.Type) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type intParser struct{}

func (p intParser) parse(value string, toType reflect.Type) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type floatParser struct{}

func (p floatParser) parse(value string, toType reflect.Type) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type timeParser struct{}

func (p timeParser) parse(value string, toType reflect.Type) (result interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type jsonParser struct{}

func (p jsonParser) parse(value string, toType reflect.Type) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sliceParser(elemParser paramParser) paramParser {
	_ = "STUB: not implemented"
	return *new(paramParser)
}

func ptrParser(elemParser paramParser) paramParser {
	_ = "STUB: not implemented"
	return *new(paramParser)
}
