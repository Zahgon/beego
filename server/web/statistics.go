package web

import (
	"sync"
	"time"
)

type Statistics struct {
	RequestURL        string
	RequestController string
	RequestNum        int64
	MinTime           time.Duration
	MaxTime           time.Duration
	TotalTime         time.Duration
}

type URLMap struct {
	lock        sync.RWMutex
	LengthLimit int
	urlmap      map[string]map[string]*Statistics
}

func (m *URLMap) AddStatistics(requestMethod, requestURL, requestController string, requesttime time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (m *URLMap) GetMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (m *URLMap) GetMapData() []map[string]interface{} { _ = "STUB: not implemented"; return nil }

var StatisticsMap *URLMap

func init() {
	StatisticsMap = &URLMap{
		urlmap: make(map[string]map[string]*Statistics),
	}
}
