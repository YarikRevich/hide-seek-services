package interfaces

import "github.com/valyala/fasthttp"

type Service interface{
	Handler(*fasthttp.RequestCtx)
	Send([]byte)
}