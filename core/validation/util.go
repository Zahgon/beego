package validation

import (
	"fmt"
	"reflect"
)

const (
	ValidTag = "valid"

	LabelTag = "label"

	wordsize = 32 << (^uint(0) >> 32 & 1)
)

var (
	funcs = make(Funcs)

	unFuncs = map[string]bool{
		"Clear":     true,
		"HasErrors": true,
		"ErrorMap":  true,
		"Error":     true,
		"apply":     true,
		"Check":     true,
		"Valid":     true,
		"NoMatch":   true,
	}

	ErrInt64On32 = fmt.Errorf("not support int64 on 32-bit platform")
)

func init() {
	v := &Validation{}
	t := reflect.TypeOf(v)
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		if !unFuncs[m.Name] {
			funcs[m.Name] = m.Func
		}
	}
}

type CustomFunc func(v *Validation, obj interface{}, key string)

func AddCustomFunc(name string, f CustomFunc) error { _ = "STUB: not implemented"; return nil }

type ValidFunc struct {
	Name   string
	Params []interface{}
}

type Funcs map[string]reflect.Value

func (f Funcs) Call(name string, params ...interface{}) (result []reflect.Value, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isStruct(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func isStructPtr(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func getValidFuncs(f reflect.StructField) (vfs []ValidFunc, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getRegFuncs(tag, key string) (vfs []ValidFunc, str string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func parseFunc(vfunc, key string, label string) (v ValidFunc, err error) {
	_ = "STUB: not implemented"
	return *new(ValidFunc), nil
}

func numIn(name string) (num int, err error) { _ = "STUB: not implemented"; return 0, nil }

func trim(name, key string, s []string) (ts []interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseParam(t reflect.Type, s string) (i interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeParam(v *Validation, obj interface{}, params []interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}
