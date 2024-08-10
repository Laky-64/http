package http

import "github.com/Laky-64/http/types"

type headersOption map[string]string

func (ct headersOption) Apply(o *types.RequestOptions) {
	o.Headers = ct
}

func Headers(headers map[string]string) RequestOption {
	return headersOption(headers)
}
