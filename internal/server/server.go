package server

import (
	"log"

	"github.com/YarikRevich/HideSeek-Services/internal/handlers"
	"github.com/YarikRevich/HideSeek-Services/internal/urls"
	bgsave "github.com/YarikRevich/HideSeek-Services/pkg/bg_save"
	"github.com/valyala/fasthttp"
)

func Init(){
	urls.PinURLToServices("localhost:8090")

	handlers.AddHandler(bgsave.UseService().Name(), bgsave.UseService().Handler())
}

func Run()chan int{
	handler := func(ctx *fasthttp.RequestCtx){
		handlers.ChooseHandler(string(ctx.Path()), ctx)
	}

	notificator := make(chan int)
	go func(){
		if err := fasthttp.ListenAndServe("localhost:8090", handler); err != nil{
			log.Fatalln(err)
		}
		notificator <- 1
	}()
	return notificator
}

func Wait(e chan int){
	for range e{
		return
	}
}