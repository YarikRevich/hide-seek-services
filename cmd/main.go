package main

import (
	"log"

	"github.com/valyala/fasthttp"

	"github.com/YarikRevich/HideSeek-Services/internal/handlers"
	"github.com/YarikRevich/HideSeek-Services/pkg/bg_save"
)

func init(){
	handlers.AddHandler("bg_save", bgsave.UseService())
}

func main(){
	handler := func(ctx *fasthttp.RequestCtx){
		handlers.ChooseHandler(string(ctx.Path()), ctx)
	}

	log.Fatalln(fasthttp.ListenAndServe(":8090", handler))
}
