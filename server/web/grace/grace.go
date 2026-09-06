package grace

import (
	"flag"
	"net/http"
	"os"
	"sync"
	"syscall"
	"time"
)

const (
	PreSignal = iota

	PostSignal

	StateInit

	StateRunning

	StateShuttingDown

	StateTerminate
)

var (
	regLock              *sync.Mutex
	runningServers       map[string]*Server
	runningServersOrder  []string
	socketPtrOffsetMap   map[string]uint
	runningServersForked bool

	DefaultReadTimeOut time.Duration

	DefaultWriteTimeOut time.Duration

	DefaultMaxHeaderBytes int

	DefaultTimeout = 60 * time.Second

	isChild     bool
	socketOrder string

	hookableSignals []os.Signal
)

func init() {
	flag.BoolVar(&isChild, "graceful", false, "listen on open fd (after forking)")
	flag.StringVar(&socketOrder, "socketorder", "", "previous initialization order - used when more than one listener was started")

	regLock = &sync.Mutex{}
	runningServers = make(map[string]*Server)
	runningServersOrder = []string{}
	socketPtrOffsetMap = make(map[string]uint)

	hookableSignals = []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
	}
}

type ServerOption func(*Server)

func WithShutdownCallback(shutdownCallback func()) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

func NewServer(addr string, handler http.Handler, opts ...ServerOption) (srv *Server) {
	_ = "STUB: not implemented"
	return nil
}

func ListenAndServe(addr string, handler http.Handler) error { _ = "STUB: not implemented"; return nil }

func ListenAndServeTLS(addr string, certFile string, keyFile string, handler http.Handler) error {
	_ = "STUB: not implemented"
	return nil
}
