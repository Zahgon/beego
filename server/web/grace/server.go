package grace

import (
	"net"
	"net/http"
	"os"
)

type Server struct {
	*http.Server
	ln                net.Listener
	SignalHooks       map[int]map[os.Signal][]func()
	sigChan           chan os.Signal
	isChild           bool
	state             uint8
	Network           string
	terminalChan      chan error
	shutdownCallbacks []func()
}

func (srv *Server) Serve() (err error) { _ = "STUB: not implemented"; return nil }

func (srv *Server) ServeWithListener(ln net.Listener) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (srv *Server) internalServe(ln net.Listener) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (srv *Server) ListenAndServe() (err error) { _ = "STUB: not implemented"; return nil }

func (srv *Server) ListenAndServeTLS(certFile, keyFile string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (srv *Server) ListenTLS(certFile string, keyFile string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func (srv *Server) ListenAndServeMutualTLS(certFile, keyFile, trustFile string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (srv *Server) ServeTLS(ln net.Listener) error { _ = "STUB: not implemented"; return nil }

func (srv *Server) ListenMutualTLS(certFile string, keyFile string, trustFile string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func (srv *Server) getListener(laddr string) (l net.Listener, err error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (srv *Server) handleSignals() { _ = "STUB: not implemented"; return }

func (srv *Server) signalHooks(ppFlag int, sig os.Signal) { _ = "STUB: not implemented"; return }

func (srv *Server) shutdown() { _ = "STUB: not implemented"; return }

func (srv *Server) fork() (err error) { _ = "STUB: not implemented"; return nil }

func (srv *Server) RegisterSignalHook(ppFlag int, sig os.Signal, f func()) (err error) {
	_ = "STUB: not implemented"
	return nil
}
