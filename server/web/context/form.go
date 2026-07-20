package context

import (
	"net/url"
	"reflect"
	"time"
)

var (
	sliceOfInts    = reflect.TypeOf([]int(nil))
	sliceOfStrings = reflect.TypeOf([]string(nil))
)

func parseFormToStruct(form url.Values, objT reflect.Type, objV reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func parseFormTime(value string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func parseFormBoolValue(value string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func formTagName(fieldT reflect.StructField) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func formFirstValue(tag string, form url.Values, fieldT reflect.StructField) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
