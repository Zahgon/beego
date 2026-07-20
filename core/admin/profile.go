package admin

import (
	"io"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

var (
	startTime = time.Now()
	pid       int
)

func init() {
	pid = os.Getpid()
}

func ProcessInput(input string, w io.Writer) { _ = "STUB: not implemented"; return }

func MemProf(w io.Writer) { _ = "STUB: not implemented"; return }

func GetCPUProfile(w io.Writer) { _ = "STUB: not implemented"; return }

func PrintGCSummary(w io.Writer) { _ = "STUB: not implemented"; return }

func printGC(memStats *runtime.MemStats, gcstats *debug.GCStats, w io.Writer) {
	_ = "STUB: not implemented"
	return
}

func avg(items []time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func toH(bytes uint64) string { _ = "STUB: not implemented"; return "" }
