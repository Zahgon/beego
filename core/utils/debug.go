package utils

import (
	"bytes"
	"reflect"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
)

type pointerInfo struct {
	prev *pointerInfo
	n    int
	addr uintptr
	pos  int
	used []int
}

func Display(data ...interface{}) { _ = "STUB: not implemented"; return }

func GetDisplayString(data ...interface{}) string { _ = "STUB: not implemented"; return "" }

func display(displayed bool, data ...interface{}) string { _ = "STUB: not implemented"; return "" }

func fomateinfo(headlen int, data ...interface{}) []byte { _ = "STUB: not implemented"; return nil }

func isSimpleType(val reflect.Value, kind reflect.Kind, pointers **pointerInfo, interfaces *[]reflect.Value) bool {
	_ = "STUB: not implemented"
	return false
}

func printKeyValue(buf *bytes.Buffer, val reflect.Value, pointers **pointerInfo, interfaces *[]reflect.Value, structFilter func(string, string) bool, formatOutput bool, indent string, level int) {
	_ = "STUB: not implemented"
	return
}

func PrintPointerInfo(buf *bytes.Buffer, headlen int, pointers *pointerInfo) {
	_ = "STUB: not implemented"
	return
}

func Stack(skip int, indent string) []byte { _ = "STUB: not implemented"; return nil }

func function(pc uintptr) []byte { _ = "STUB: not implemented"; return nil }
