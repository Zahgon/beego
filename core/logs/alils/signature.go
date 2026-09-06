package alils

import (
	"time"
)

var gmtLoc = time.FixedZone("GMT", 0)

func nowRFC1123() string { _ = "STUB: not implemented"; return "" }

func signature(project *LogProject, method, uri string,
	headers map[string]string) (digest string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getHeaderSafe(headers map[string]string, key string) string {
	_ = "STUB: not implemented"
	return ""
}

func buildCanonicalHeaders(headers map[string]string) string { _ = "STUB: not implemented"; return "" }

func buildCanonicalResource(uri string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func buildSignString(method, contentMD5, contentType, date, canoHeaders, canoResource string) string {
	_ = "STUB: not implemented"
	return ""
}

func calculateHmacSha1(signStr, secret string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
