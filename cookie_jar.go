package http

import (
	"github.com/Laky-64/http/types"
	"net/http"
)

type cookieJarOption struct {
	cookieJar http.CookieJar
}

func (ct cookieJarOption) Apply(o *types.RequestOptions) {
	o.CookieJar = ct.cookieJar
}

func CookieJar(cookieJar http.CookieJar) RequestOption {
	return cookieJarOption{cookieJar}
}
