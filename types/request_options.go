package types

import (
	"io"
	"time"
)

type RequestOptions struct {
	Retries        int
	Timeout        time.Duration
	Method         string
	BearerToken    string
	Body           []byte
	Headers        map[string]string
	Cookies        map[string]string
	MultiPart      *MultiPartInfo
	OverloadReader func(r io.Reader) io.Reader
	Proxy          string
}
