package handlers

import (
	"github.com/YarikRevich/HideSeek-Services/tools"
	"github.com/valyala/fasthttp"
)

var (
	handlers = map[string]fasthttp.RequestHandler{}
)

func ChooseHandler(path string, ctx *fasthttp.RequestCtx){
	for k, v := range handlers{
		if tools.IsService(path, k){
			v(ctx)
		}
	}
} 

func AddHandler(serviceName string, handler fasthttp.RequestHandler){
	handlers[serviceName] = handler
}