package bean

import (
	"reflect"
)

func Mock(v interface{}) (err error) { _ = "STUB: not implemented"; return nil }

func mock(pv reflect.Value) (err error) { _ = "STUB: not implemented"; return nil }

func mockSlice(tagValue string, pvv reflect.Value) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func mockBool(tagValue string, pvv reflect.Value) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func mockPtr(pvv reflect.Value, ptt reflect.Type) (err error) {
	_ = "STUB: not implemented"
	return nil
}
