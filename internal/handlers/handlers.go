package handlers

import (
	"github.com/YarikRevich/HideSeek-Services/pkg/interfaces"
	"github.com/YarikRevich/HideSeek-Services/tools"
	"github.com/valyala/fasthttp"
)

var (
	handlers map[string]interfaces.Service
)

func ChooseHandler(path string, ctx *fasthttp.RequestCtx){
	for k, v := range handlers{
		if tools.IsService(path, k){
			v.Handler(ctx)
		}
	}
} 

func AddHandler(serviceName string, service interfaces.Service){
	handlers[serviceName] = service
}