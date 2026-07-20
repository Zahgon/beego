package web

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"sync"
)

var (
	beegoTplFuncMap           = make(template.FuncMap)
	beeViewPathTemplateLocked = false

	beeViewPathTemplates = make(map[string]map[string]*template.Template)
	templatesLock        sync.RWMutex

	beeTemplateExt = []string{"tpl", "html", "gohtml"}

	beeTemplateEngines = map[string]templatePreProcessor{}
	beeTemplateFS      = defaultFSFunc
)

func ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func ExecuteViewPathTemplate(wr io.Writer, name string, viewPath string, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	beegoTplFuncMap["dateformat"] = DateFormat
	beegoTplFuncMap["date"] = Date
	beegoTplFuncMap["compare"] = Compare
	beegoTplFuncMap["compare_not"] = CompareNot
	beegoTplFuncMap["not_nil"] = NotNil
	beegoTplFuncMap["not_null"] = NotNil
	beegoTplFuncMap["substr"] = Substr
	beegoTplFuncMap["html2str"] = HTML2str
	beegoTplFuncMap["str2html"] = Str2html
	beegoTplFuncMap["htmlquote"] = Htmlquote
	beegoTplFuncMap["htmlunquote"] = Htmlunquote
	beegoTplFuncMap["renderform"] = RenderForm
	beegoTplFuncMap["assets_js"] = AssetsJs
	beegoTplFuncMap["assets_css"] = AssetsCSS
	beegoTplFuncMap["config"] = GetConfig
	beegoTplFuncMap["map_get"] = MapGet

	beegoTplFuncMap["eq"] = eq
	beegoTplFuncMap["ge"] = ge
	beegoTplFuncMap["gt"] = gt
	beegoTplFuncMap["le"] = le
	beegoTplFuncMap["lt"] = lt
	beegoTplFuncMap["ne"] = ne

	beegoTplFuncMap["urlfor"] = URLFor
}

func AddFuncMap(key string, fn interface{}) error { _ = "STUB: not implemented"; return nil }

type templatePreProcessor func(root, path string, funcs template.FuncMap) (*template.Template, error)

type templateFile struct {
	root  string
	files map[string][]string
}

func (tf *templateFile) visit(paths string, f os.FileInfo, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func HasTemplateExt(paths string) bool { _ = "STUB: not implemented"; return false }

func AddTemplateExt(ext string) { _ = "STUB: not implemented"; return }

func AddViewPath(viewPath string) error { _ = "STUB: not implemented"; return nil }

func lockViewPaths() { _ = "STUB: not implemented"; return }

func BuildTemplate(dir string, files ...string) error { _ = "STUB: not implemented"; return nil }

func getTplDeep(root string, fs http.FileSystem, file string, parent string, t *template.Template) (*template.Template, [][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func getTemplate(root string, fs http.FileSystem, file string, others ...string) (t *template.Template, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func _getTemplate(t0 *template.Template, root string, fs http.FileSystem, subMods [][]string, others ...string) (t *template.Template, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type templateFSFunc func() http.FileSystem

func defaultFSFunc() http.FileSystem { _ = "STUB: not implemented"; return *new(http.FileSystem) }

func SetTemplateFSFunc(fnt templateFSFunc) { _ = "STUB: not implemented"; return }

func SetViewsPath(path string) *HttpServer { _ = "STUB: not implemented"; return nil }

func SetStaticPath(url string, path string) *HttpServer { _ = "STUB: not implemented"; return nil }

func DelStaticPath(url string) *HttpServer { _ = "STUB: not implemented"; return nil }

func AddTemplateEngine(extension string, fn templatePreProcessor) *HttpServer {
	_ = "STUB: not implemented"
	return nil
}
