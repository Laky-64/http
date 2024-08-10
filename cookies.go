package http

import "github.com/Laky-64/http/types"

type cookiesOption map[string]string

func (ct cookiesOption) Apply(o *types.RequestOptions) {
	o.Cookies = ct
}

func Cookies(cookies map[string]string) RequestOption {
	return cookiesOption(cookies)
}
