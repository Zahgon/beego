package logs

import (
	"os"
	"sync"
	"time"
)

type fileLogWriter struct {
	sync.RWMutex

	Rotate bool `json:"rotate"`
	Daily  bool `json:"daily"`
	Hourly bool `json:"hourly"`

	Filename   string `json:"filename"`
	fileWriter *os.File

	MaxLines         int `json:"maxlines"`
	maxLinesCurLines int

	MaxFiles         int `json:"maxfiles"`
	MaxFilesCurFiles int

	MaxSize        int `json:"maxsize"`
	maxSizeCurSize int

	MaxDays       int64 `json:"maxdays"`
	dailyOpenDate int
	dailyOpenTime time.Time

	MaxHours       int64 `json:"maxhours"`
	hourlyOpenDate int
	hourlyOpenTime time.Time

	Level int `json:"level"`

	Perm string `json:"perm"`

	DirPerm string `json:"dirperm"`

	RotatePerm string `json:"rotateperm"`

	fileNameOnly, suffix string

	logFormatter LogFormatter
	Formatter    string `json:"formatter"`
}

func newFileWriter() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (*fileLogWriter) Format(lm *LogMsg) string { _ = "STUB: not implemented"; return "" }

func (w *fileLogWriter) SetFormatter(f LogFormatter) { _ = "STUB: not implemented"; return }

func (w *fileLogWriter) Init(config string) error { _ = "STUB: not implemented"; return nil }

func (w *fileLogWriter) startLogger() error { _ = "STUB: not implemented"; return nil }

func (w *fileLogWriter) needRotateDaily(day int) bool { _ = "STUB: not implemented"; return false }

func (w *fileLogWriter) needRotateHourly(hour int) bool { _ = "STUB: not implemented"; return false }

func (w *fileLogWriter) WriteMsg(lm *LogMsg) error { _ = "STUB: not implemented"; return nil }

func (w *fileLogWriter) createLogFile() (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *fileLogWriter) initFd() error { _ = "STUB: not implemented"; return nil }

func (w *fileLogWriter) dailyRotate(openTime time.Time) { _ = "STUB: not implemented"; return }

func (w *fileLogWriter) hourlyRotate(openTime time.Time) { _ = "STUB: not implemented"; return }

func (w *fileLogWriter) lines() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *fileLogWriter) doRotate(logTime time.Time) error { _ = "STUB: not implemented"; return nil }

func (w *fileLogWriter) deleteOldLog() { _ = "STUB: not implemented"; return }

func (w *fileLogWriter) Destroy() { _ = "STUB: not implemented"; return }

func (w *fileLogWriter) Flush() { _ = "STUB: not implemented"; return }

func init() {
	Register(AdapterFile, newFileWriter)
}
