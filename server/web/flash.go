package web

type FlashData struct {
	Data map[string]string
}

func NewFlash() *FlashData { _ = "STUB: not implemented"; return nil }

func (fd *FlashData) Set(key string, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (fd *FlashData) Success(msg string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (fd *FlashData) Notice(msg string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (fd *FlashData) Warning(msg string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (fd *FlashData) Error(msg string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (fd *FlashData) Store(c *Controller) { _ = "STUB: not implemented"; return }

func ReadFromRequest(c *Controller) *FlashData { _ = "STUB: not implemented"; return nil }
