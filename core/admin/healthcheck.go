package admin

var AdminCheckList map[string]HealthChecker

type HealthChecker interface {
	Check() error
}

func AddHealthCheck(name string, hc HealthChecker) { _ = "STUB: not implemented"; return }

func init() {
	AdminCheckList = make(map[string]HealthChecker)
}
