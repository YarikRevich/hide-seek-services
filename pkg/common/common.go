package common

import "github.com/valyala/fasthttp"

type Service interface {
	Handler() fasthttp.RequestHandler
	SetURL(url string)
	Send([]byte)
	SendAsync([]byte)

	Name() string
}
