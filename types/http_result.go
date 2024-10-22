package types

import "net/http"

type HTTPResult struct {
	StatusCode int
	Headers    map[string][]string
	Cookies    []*http.Cookie
	Body       []byte
}

func (r *HTTPResult) String() string {
	return string(r.Body)
}
