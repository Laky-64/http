package http

import "github.com/Laky-64/http/types"

type retriesOption int

func (ct retriesOption) Apply(o *types.RequestOptions) {
	o.Retries = int(ct)
}

func Retries(count int) RequestOption {
	return retriesOption(count)
}
