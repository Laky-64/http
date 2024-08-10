package http

import "github.com/Laky-64/http/types"

type bearerTokenOption string

func (ct bearerTokenOption) Apply(o *types.RequestOptions) {
	o.BearerToken = string(ct)
}

func BearerToken(method string) RequestOption {
	return bearerTokenOption(method)
}
