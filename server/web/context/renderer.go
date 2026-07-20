package context

type Renderer interface {
	Render(ctx *Context)
}

type rendererFunc func(ctx *Context)

func (f rendererFunc) Render(ctx *Context) { _ = "STUB: not implemented"; return }
