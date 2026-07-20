package mock

import (
	"context"

	"github.com/beego/beego/v2/client/httplib"
)

type RequestCondition interface {
	Match(ctx context.Context, req *httplib.BeegoHTTPRequest) bool
}

type SimpleCondition struct {
	pathReg string
	path    string
	method  string
	query   map[string]string
	header  map[string]string
	body    map[string]interface{}
}

func NewSimpleCondition(path string, opts ...simpleConditionOption) *SimpleCondition {
	_ = "STUB: not implemented"
	return nil
}

func (sc *SimpleCondition) Match(ctx context.Context, req *httplib.BeegoHTTPRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func (sc *SimpleCondition) matchPath(ctx context.Context, req *httplib.BeegoHTTPRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func (sc *SimpleCondition) matchPathReg(ctx context.Context, req *httplib.BeegoHTTPRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func (sc *SimpleCondition) matchQuery(ctx context.Context, req *httplib.BeegoHTTPRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func (sc *SimpleCondition) matchHeader(ctx context.Context, req *httplib.BeegoHTTPRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func (sc *SimpleCondition) matchBodyFields(ctx context.Context, req *httplib.BeegoHTTPRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func (sc *SimpleCondition) matchMethod(ctx context.Context, req *httplib.BeegoHTTPRequest) bool {
	_ = "STUB: not implemented"
	return false
}

type simpleConditionOption func(sc *SimpleCondition)

func WithPathReg(pathReg string) simpleConditionOption {
	_ = "STUB: not implemented"
	return *new(simpleConditionOption)
}

func WithQuery(key, value string) simpleConditionOption {
	_ = "STUB: not implemented"
	return *new(simpleConditionOption)
}

func WithHeader(key, value string) simpleConditionOption {
	_ = "STUB: not implemented"
	return *new(simpleConditionOption)
}

func WithJsonBodyFields(field string, value interface{}) simpleConditionOption {
	_ = "STUB: not implemented"
	return *new(simpleConditionOption)
}

func WithMethod(method string) simpleConditionOption {
	_ = "STUB: not implemented"
	return *new(simpleConditionOption)
}
