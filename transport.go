package http

import (
	"github.com/Laky-64/http/types"
	"net/http"
)

type transportOption struct {
	transport *http.Transport
}

func (t transportOption) Apply(o *types.RequestOptions) {
	o.Transport = t.transport
}

func Transport(transport *http.Transport) RequestOption {
	return transportOption{transport}
}
