package http

import "github.com/Laky-64/http/types"

type RequestOption interface {
	Apply(o *types.RequestOptions)
}
