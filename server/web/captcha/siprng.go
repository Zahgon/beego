package captcha

import (
	"sync"
)

type siprng struct {
	mu          sync.Mutex
	k0, k1, ctr uint64
}

func siphash(k0, k1, m uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func (p *siprng) rekey() { _ = "STUB: not implemented"; return }

func (p *siprng) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *siprng) Int63() int64 { _ = "STUB: not implemented"; return 0 }

func (p *siprng) Uint32() uint32 { _ = "STUB: not implemented"; return 0 }

func (p *siprng) Int31() int32 { _ = "STUB: not implemented"; return 0 }

func (p *siprng) Intn(n int) int { _ = "STUB: not implemented"; return 0 }

func (p *siprng) Int63n(n int64) int64 { _ = "STUB: not implemented"; return 0 }

func (p *siprng) Int31n(n int32) int32 { _ = "STUB: not implemented"; return 0 }

func (p *siprng) Float64() float64 { _ = "STUB: not implemented"; return 0 }
