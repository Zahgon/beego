package logs

import (
	"time"
)

type LogMsg struct {
	Level               int
	Msg                 string
	When                time.Time
	FilePath            string
	LineNumber          int
	Args                []interface{}
	Prefix              string
	enableFullFilePath  bool
	enableFuncCallDepth bool
}

func (lm *LogMsg) OldStyleFormat() string { _ = "STUB: not implemented"; return "" }
