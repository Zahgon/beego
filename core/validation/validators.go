package validation

import (
	"regexp"
	"sync"
)

var CanSkipFuncs = map[string]struct{}{
	"Email":   {},
	"IP":      {},
	"Mobile":  {},
	"Tel":     {},
	"Phone":   {},
	"ZipCode": {},
}

var MessageTmpls = map[string]string{
	"Required":     "Can not be empty",
	"Min":          "Minimum is %d",
	"Max":          "Maximum is %d",
	"Range":        "Range is %d to %d",
	"MinSize":      "Minimum size is %d",
	"MaxSize":      "Maximum size is %d",
	"Length":       "Required length is %d",
	"Alpha":        "Must be valid alpha characters",
	"Numeric":      "Must be valid numeric characters",
	"AlphaNumeric": "Must be valid alpha or numeric characters",
	"Match":        "Must match %s",
	"NoMatch":      "Must not match %s",
	"AlphaDash":    "Must be valid alpha or numeric or dash(-_) characters",
	"Email":        "Must be a valid email address",
	"IP":           "Must be a valid ip address",
	"Base64":       "Must be valid base64 characters",
	"Mobile":       "Must be valid mobile number",
	"Tel":          "Must be valid telephone number",
	"Phone":        "Must be valid telephone or mobile phone number",
	"ZipCode":      "Must be valid zipcode",
	"Enum":         "Must be a string value in \"%s\"",
}

var once sync.Once

func SetDefaultMessage(msg map[string]string) { _ = "STUB: not implemented"; return }

type Validator interface {
	IsSatisfied(interface{}) bool
	DefaultMessage() string
	GetKey() string
	GetLimitValue() interface{}
}

type Required struct {
	Key string
}

func (r Required) IsSatisfied(obj interface{}) bool { _ = "STUB: not implemented"; return false }

func (r Required) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (r Required) GetKey() string { _ = "STUB: not implemented"; return "" }

func (r Required) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

type Min struct {
	Min int
	Key string
}

func (m Min) IsSatisfied(obj interface{}) bool { _ = "STUB: not implemented"; return false }

func (m Min) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (m Min) GetKey() string { _ = "STUB: not implemented"; return "" }

func (m Min) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

type Max struct {
	Max int
	Key string
}

func (m Max) IsSatisfied(obj interface{}) bool { _ = "STUB: not implemented"; return false }

func (m Max) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (m Max) GetKey() string { _ = "STUB: not implemented"; return "" }

func (m Max) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

type Range struct {
	Min
	Max
	Key string
}

func (r Range) IsSatisfied(obj interface{}) bool { _ = "STUB: not implemented"; return false }

func (r Range) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (r Range) GetKey() string { _ = "STUB: not implemented"; return "" }

func (r Range) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

type MinSize struct {
	Min int
	Key string
}

func (m MinSize) IsSatisfied(obj interface{}) bool { _ = "STUB: not implemented"; return false }

func (m MinSize) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (m MinSize) GetKey() string { _ = "STUB: not implemented"; return "" }

func (m MinSize) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

type MaxSize struct {
	Max int
	Key string
}

func (m MaxSize) IsSatisfied(obj interface{}) bool { _ = "STUB: not implemented"; return false }

func (m MaxSize) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (m MaxSize) GetKey() string { _ = "STUB: not implemented"; return "" }

func (m MaxSize) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

type Length struct {
	N   int
	Key string
}

func (l Length) IsSatisfied(obj interface{}) bool { _ = "STUB: not implemented"; return false }

func (l Length) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (l Length) GetKey() string { _ = "STUB: not implemented"; return "" }

func (l Length) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

type Alpha struct {
	Key string
}

func (a Alpha) IsSatisfied(obj interface{}) bool { _ = "STUB: not implemented"; return false }

func (a Alpha) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (a Alpha) GetKey() string { _ = "STUB: not implemented"; return "" }

func (a Alpha) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

type Numeric struct {
	Key string
}

func (n Numeric) IsSatisfied(obj interface{}) bool { _ = "STUB: not implemented"; return false }

func (n Numeric) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (n Numeric) GetKey() string { _ = "STUB: not implemented"; return "" }

func (n Numeric) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

type AlphaNumeric struct {
	Key string
}

func (a AlphaNumeric) IsSatisfied(obj interface{}) bool { _ = "STUB: not implemented"; return false }

func (a AlphaNumeric) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (a AlphaNumeric) GetKey() string { _ = "STUB: not implemented"; return "" }

func (a AlphaNumeric) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

type Match struct {
	Regexp *regexp.Regexp
	Key    string
}

func (m Match) IsSatisfied(obj interface{}) bool { _ = "STUB: not implemented"; return false }

func (m Match) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (m Match) GetKey() string { _ = "STUB: not implemented"; return "" }

func (m Match) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

type NoMatch struct {
	Match
	Key string
}

func (n NoMatch) IsSatisfied(obj interface{}) bool { _ = "STUB: not implemented"; return false }

func (n NoMatch) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (n NoMatch) GetKey() string { _ = "STUB: not implemented"; return "" }

func (n NoMatch) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

var alphaDashPattern = regexp.MustCompile(`[^\d\w-_]`)

type AlphaDash struct {
	NoMatch
	Key string
}

func (a AlphaDash) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (a AlphaDash) GetKey() string { _ = "STUB: not implemented"; return "" }

func (a AlphaDash) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

var emailPattern = regexp.MustCompile(`^[\w!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\.[\w!#$%&'*+/=?^_` + "`" + `{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[a-zA-Z0-9](?:[\w-]*[\w])?$`)

type Email struct {
	Match
	Key string
}

func (e Email) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (e Email) GetKey() string { _ = "STUB: not implemented"; return "" }

func (e Email) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

var ipPattern = regexp.MustCompile(`^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$`)

type IP struct {
	Match
	Key string
}

func (i IP) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (i IP) GetKey() string { _ = "STUB: not implemented"; return "" }

func (i IP) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

var base64Pattern = regexp.MustCompile(`^(?:[A-Za-z0-99+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$`)

type Base64 struct {
	Match
	Key string
}

func (b Base64) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (b Base64) GetKey() string { _ = "STUB: not implemented"; return "" }

func (b Base64) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

var mobilePattern = regexp.MustCompile(`^((\+86)|(86))?1([356789][0-9]|4[579]|6[67]|7[0135678]|9[189])[0-9]{8}$`)

type Mobile struct {
	Match
	Key string
}

func (m Mobile) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (m Mobile) GetKey() string { _ = "STUB: not implemented"; return "" }

func (m Mobile) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

var telPattern = regexp.MustCompile(`^(0\d{2,3}(\-)?)?\d{7,8}$`)

type Tel struct {
	Match
	Key string
}

func (t Tel) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (t Tel) GetKey() string { _ = "STUB: not implemented"; return "" }

func (t Tel) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

type Phone struct {
	Mobile
	Tel
	Key string
}

func (p Phone) IsSatisfied(obj interface{}) bool { _ = "STUB: not implemented"; return false }

func (p Phone) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (p Phone) GetKey() string { _ = "STUB: not implemented"; return "" }

func (p Phone) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

var zipCodePattern = regexp.MustCompile(`^[1-9]\d{5}$`)

type ZipCode struct {
	Match
	Key string
}

func (z ZipCode) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (z ZipCode) GetKey() string { _ = "STUB: not implemented"; return "" }

func (z ZipCode) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }

type Enum struct {
	Rules string
	Key   string
}

func (e Enum) IsSatisfied(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (e Enum) DefaultMessage() string { _ = "STUB: not implemented"; return "" }

func (e Enum) GetKey() string { _ = "STUB: not implemented"; return "" }

func (Enum) GetLimitValue() interface{} { _ = "STUB: not implemented"; return nil }
