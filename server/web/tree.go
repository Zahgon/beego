package web

import (
	"regexp"

	"github.com/beego/beego/v2/server/web/context"
)

var allowSuffixExt = []string{".json", ".xml", ".html"}

type Tree struct {
	prefix string

	fixrouters []*Tree

	wildcard *Tree

	leaves []*leafInfo
}

func NewTree() *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) AddTree(prefix string, tree *Tree) { _ = "STUB: not implemented"; return }

func (t *Tree) addtree(segments []string, tree *Tree, wildcards []string, reg string) {
	_ = "STUB: not implemented"
	return
}

func filterTreeWithPrefix(t *Tree, wildcards []string, reg string) {
	_ = "STUB: not implemented"
	return
}

func (t *Tree) AddRouter(pattern string, runObject interface{}) { _ = "STUB: not implemented"; return }

func (t *Tree) addseg(segments []string, route interface{}, wildcards []string, reg string) {
	_ = "STUB: not implemented"
	return
}

func (t *Tree) Match(pattern string, ctx *context.Context) (runObject interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tree) match(treePattern string, pattern string, wildcardValues []string, ctx *context.Context) (runObject interface{}) {
	_ = "STUB: not implemented"
	return nil
}

type leafInfo struct {
	wildcards []string

	regexps *regexp.Regexp

	runObject interface{}
}

func (leaf *leafInfo) match(treePattern string, wildcardValues []string, ctx *context.Context) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func splitPath(key string) []string { _ = "STUB: not implemented"; return nil }

func splitSegment(key string) (bool, []string, string) {
	_ = "STUB: not implemented"
	return false, nil, ""
}
