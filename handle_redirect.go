package http

import (
	"github.com/Laky-64/http/types"
	"net/http"
)

type handleRedirect func(req *http.Request, via []*http.Request) error

func (ct handleRedirect) Apply(o *types.RequestOptions) {
	o.HandleRedirect = ct
}

func HandleRedirect(handle func(req *http.Request, via []*http.Request) error) RequestOption {
	return handleRedirect(handle)
}
