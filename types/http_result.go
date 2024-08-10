package types

type HTTPResult struct {
	StatusCode int
	Body       []byte
}

func (r *HTTPResult) String() string {
	return string(r.Body)
}
