package pagination

import (
	"net/http"
)

type Paginator struct {
	Request     *http.Request
	PerPageNums int
	MaxPages    int

	nums      int64
	pageRange []int
	pageNums  int
	page      int
}

func (p *Paginator) PageNums() int { _ = "STUB: not implemented"; return 0 }

func (p *Paginator) Nums() int64 { _ = "STUB: not implemented"; return 0 }

func (p *Paginator) SetNums(nums interface{}) { _ = "STUB: not implemented"; return }

func (p *Paginator) Page() int { _ = "STUB: not implemented"; return 0 }

func (p *Paginator) Pages() []int { _ = "STUB: not implemented"; return nil }

func (p *Paginator) PageLink(page int) string { _ = "STUB: not implemented"; return "" }

func (p *Paginator) PageLinkPrev() (link string) { _ = "STUB: not implemented"; return "" }

func (p *Paginator) PageLinkNext() (link string) { _ = "STUB: not implemented"; return "" }

func (p *Paginator) PageLinkFirst() (link string) { _ = "STUB: not implemented"; return "" }

func (p *Paginator) PageLinkLast() (link string) { _ = "STUB: not implemented"; return "" }

func (p *Paginator) HasPrev() bool { _ = "STUB: not implemented"; return false }

func (p *Paginator) HasNext() bool { _ = "STUB: not implemented"; return false }

func (p *Paginator) IsActive(page int) bool { _ = "STUB: not implemented"; return false }

func (p *Paginator) Offset() int { _ = "STUB: not implemented"; return 0 }

func (p *Paginator) HasPages() bool { _ = "STUB: not implemented"; return false }

func NewPaginator(req *http.Request, per int, nums interface{}) *Paginator {
	_ = "STUB: not implemented"
	return nil
}
