package web

import (
	"net/http"
)

type adminController struct {
	Controller
	servers []*HttpServer
}

func (a *adminController) registerHttpServer(svr *HttpServer) { _ = "STUB: not implemented"; return }

func (a *adminController) ProfIndex() { _ = "STUB: not implemented"; return }

func (a *adminController) PrometheusMetrics() { _ = "STUB: not implemented"; return }

func (a *adminController) TaskStatus() { _ = "STUB: not implemented"; return }

func (a *adminController) AdminIndex() { _ = "STUB: not implemented"; return }

func (a *adminController) Healthcheck() { _ = "STUB: not implemented"; return }

func heathCheck(rw http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func (a *adminController) QpsIndex() { _ = "STUB: not implemented"; return }

func (a *adminController) ListConf() { _ = "STUB: not implemented"; return }

func writeTemplate(rw http.ResponseWriter, data map[interface{}]interface{}, tpls ...string) {
	_ = "STUB: not implemented"
	return
}

func buildHealthCheckResponseList(healthCheckResults *[][]string) []map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func PrintTree() M { _ = "STUB: not implemented"; return *new(M) }
