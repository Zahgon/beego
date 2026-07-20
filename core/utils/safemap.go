package utils

import (
	"sync"
)

type BeeMap struct {
	lock *sync.RWMutex
	bm   map[interface{}]interface{}
}

func NewBeeMap() *BeeMap { _ = "STUB: not implemented"; return nil }

func (m *BeeMap) Get(k interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func (m *BeeMap) Set(k interface{}, v interface{}) bool { _ = "STUB: not implemented"; return false }

func (m *BeeMap) Check(k interface{}) bool { _ = "STUB: not implemented"; return false }

func (m *BeeMap) Delete(k interface{}) { _ = "STUB: not implemented"; return }

func (m *BeeMap) Items() map[interface{}]interface{} { _ = "STUB: not implemented"; return nil }

func (m *BeeMap) Count() int { _ = "STUB: not implemented"; return 0 }
