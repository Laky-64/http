package types

import (
	"io"
	"net/http"
	"time"
)

type RequestOptions struct {
	Retries        int
	Timeout        time.Duration
	Method         string
	BearerToken    string
	Body           []byte
	Headers        map[string]string
	HandleRedirect func(req *http.Request, via []*http.Request) error
	Cookies        map[string]string
	MultiPart      *MultiPartInfo
	OverloadReader func(r io.Reader) io.Reader
	Proxy          string
	Transport      *http.Transport
}
