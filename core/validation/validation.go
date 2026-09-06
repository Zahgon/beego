package validation

import (
	"regexp"
)

type ValidFormer interface {
	Valid(*Validation)
}

type Error struct {
	Message, Key, Name, Field, Tmpl, Label string
	Value                                  interface{}
	LimitValue                             interface{}
}

func (e *Error) String() string { _ = "STUB: not implemented"; return "" }

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

type Result struct {
	Error *Error
	Ok    bool
}

func (r *Result) Key(key string) *Result { _ = "STUB: not implemented"; return nil }

func (r *Result) Message(message string, args ...interface{}) *Result {
	_ = "STUB: not implemented"
	return nil
}

type Validation struct {
	RequiredFirst bool

	Errors    []*Error
	ErrorsMap map[string][]*Error
}

func (v *Validation) Clear() { _ = "STUB: not implemented"; return }

func (v *Validation) HasErrors() bool { _ = "STUB: not implemented"; return false }

func (v *Validation) ErrorMap() map[string][]*Error { _ = "STUB: not implemented"; return nil }

func (v *Validation) Error(message string, args ...interface{}) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) Required(obj interface{}, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) Min(obj interface{}, min int, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) Max(obj interface{}, max int, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) Range(obj interface{}, min, max int, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) MinSize(obj interface{}, min int, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) MaxSize(obj interface{}, max int, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) Length(obj interface{}, n int, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) Alpha(obj interface{}, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) Numeric(obj interface{}, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) AlphaNumeric(obj interface{}, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) Match(obj interface{}, regex *regexp.Regexp, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) NoMatch(obj interface{}, regex *regexp.Regexp, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) AlphaDash(obj interface{}, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) Email(obj interface{}, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) IP(obj interface{}, key string) *Result { _ = "STUB: not implemented"; return nil }

func (v *Validation) Base64(obj interface{}, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) Mobile(obj interface{}, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) Tel(obj interface{}, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) Phone(obj interface{}, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) ZipCode(obj interface{}, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) Enum(obj interface{}, vals string, key string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) apply(chk Validator, obj interface{}) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) AddError(key, message string) { _ = "STUB: not implemented"; return }

func (v *Validation) setError(err *Error) { _ = "STUB: not implemented"; return }

func (v *Validation) SetError(fieldName string, errMsg string) *Error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) Check(obj interface{}, checks ...Validator) *Result {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) Valid(obj interface{}) (b bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (v *Validation) RecursiveValid(objc interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (v *Validation) CanSkipAlso(skipFunc string) { _ = "STUB: not implemented"; return }
