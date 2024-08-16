package http

import "github.com/Laky-64/http/types"

type proxyUri string

func (ct proxyUri) Apply(o *types.RequestOptions) {
	o.Proxy = string(ct)
}

func Proxy(url string) RequestOption {
	return proxyUri(url)
}
