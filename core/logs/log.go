package logs

import (
	"log"
	"sync"
)

const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

const levelLoggerImpl = -1

const (
	AdapterConsole   = "console"
	AdapterFile      = "file"
	AdapterMultiFile = "multifile"
	AdapterMail      = "smtp"
	AdapterConn      = "conn"
	AdapterEs        = "es"
	AdapterJianLiao  = "jianliao"
	AdapterSlack     = "slack"
	AdapterAliLS     = "alils"
)

const (
	LevelInfo  = LevelInformational
	LevelTrace = LevelDebug
	LevelWarn  = LevelWarning
)

type newLoggerFunc func() Logger

type Logger interface {
	Init(config string) error
	WriteMsg(lm *LogMsg) error
	Destroy()
	Flush()
	SetFormatter(f LogFormatter)
}

var (
	adapters    = make(map[string]newLoggerFunc)
	levelPrefix = [LevelDebug + 1]string{"[M]", "[A]", "[C]", "[E]", "[W]", "[N]", "[I]", "[D]"}
)

func Register(name string, log newLoggerFunc) { _ = "STUB: not implemented"; return }

type BeeLogger struct {
	lock                sync.Mutex
	init                bool
	enableFuncCallDepth bool
	enableFullFilePath  bool
	asynchronous        bool

	logWithNonBlocking  bool
	wg                  sync.WaitGroup
	level               int
	loggerFuncCallDepth int
	prefix              string
	msgChanLen          int64
	msgChan             chan *LogMsg
	closeChan           chan struct{}
	flushChan           chan struct{}
	outputs             []*nameLogger
	globalFormatter     string
}

const defaultAsyncMsgLen = 1e3

type nameLogger struct {
	Logger
	name string
}

var logMsgPool *sync.Pool

func NewLogger(channelLens ...int64) *BeeLogger { _ = "STUB: not implemented"; return nil }

func (bl *BeeLogger) Async(msgLen ...int64) *BeeLogger { _ = "STUB: not implemented"; return nil }

func (bl *BeeLogger) AsyncNonBlockWrite() *BeeLogger { _ = "STUB: not implemented"; return nil }

func (bl *BeeLogger) setLogger(adapterName string, configs ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (bl *BeeLogger) SetLogger(adapterName string, configs ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (bl *BeeLogger) DelLogger(adapterName string) error { _ = "STUB: not implemented"; return nil }

func (bl *BeeLogger) writeToLoggers(lm *LogMsg) { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (bl *BeeLogger) writeMsg(lm *LogMsg) error { _ = "STUB: not implemented"; return nil }

func (bl *BeeLogger) SetLevel(l int) { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) GetLevel() int { _ = "STUB: not implemented"; return 0 }

func (bl *BeeLogger) SetLogFuncCallDepth(d int) { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) GetLogFuncCallDepth() int { _ = "STUB: not implemented"; return 0 }

func (bl *BeeLogger) EnableFuncCallDepth(b bool) { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) SetPrefix(s string) { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) startLogger() { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) setGlobalFormatter(fmtter string) error { _ = "STUB: not implemented"; return nil }

func SetGlobalFormatter(fmtter string) error { _ = "STUB: not implemented"; return nil }

func (bl *BeeLogger) Emergency(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) Alert(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) Critical(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) Error(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) Warning(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) Notice(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) Informational(format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (bl *BeeLogger) Debug(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) Warn(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) Info(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) Trace(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) Flush() { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) Close() { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) Reset() { _ = "STUB: not implemented"; return }

func (bl *BeeLogger) flush() { _ = "STUB: not implemented"; return }

var beeLogger = NewLogger()

func GetBeeLogger() *BeeLogger { _ = "STUB: not implemented"; return nil }

var beeLoggerMap = struct {
	sync.RWMutex
	logs map[string]*log.Logger
}{
	logs: map[string]*log.Logger{},
}

func GetLogger(prefixes ...string) *log.Logger { _ = "STUB: not implemented"; return nil }

func EnableFullFilePath(b bool) { _ = "STUB: not implemented"; return }

func Reset() { _ = "STUB: not implemented"; return }

func Async(msgLen ...int64) *BeeLogger { _ = "STUB: not implemented"; return nil }

func SetLevel(l int) { _ = "STUB: not implemented"; return }

func SetPrefix(s string) { _ = "STUB: not implemented"; return }

func EnableFuncCallDepth(b bool) { _ = "STUB: not implemented"; return }

func SetLogFuncCall(b bool) { _ = "STUB: not implemented"; return }

func SetLogFuncCallDepth(d int) { _ = "STUB: not implemented"; return }

func SetLogger(adapter string, config ...string) error { _ = "STUB: not implemented"; return nil }

func Emergency(f interface{}, v ...interface{}) { _ = "STUB: not implemented"; return }

func Alert(f interface{}, v ...interface{}) { _ = "STUB: not implemented"; return }

func Critical(f interface{}, v ...interface{}) { _ = "STUB: not implemented"; return }

func Error(f interface{}, v ...interface{}) { _ = "STUB: not implemented"; return }

func Warning(f interface{}, v ...interface{}) { _ = "STUB: not implemented"; return }

func Warn(f interface{}, v ...interface{}) { _ = "STUB: not implemented"; return }

func Notice(f interface{}, v ...interface{}) { _ = "STUB: not implemented"; return }

func Informational(f interface{}, v ...interface{}) { _ = "STUB: not implemented"; return }

func Info(f interface{}, v ...interface{}) { _ = "STUB: not implemented"; return }

func Debug(f interface{}, v ...interface{}) { _ = "STUB: not implemented"; return }

func Trace(f interface{}, v ...interface{}) { _ = "STUB: not implemented"; return }

func formatPattern(f interface{}, v ...interface{}) string { _ = "STUB: not implemented"; return "" }
