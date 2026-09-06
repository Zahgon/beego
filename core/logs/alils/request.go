package alils

import (
	"net/http"
)

func request(project *LogProject, method, uri string, headers map[string]string,
	body []byte) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
