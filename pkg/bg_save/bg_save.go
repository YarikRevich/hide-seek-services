package bgsave

import (
	"log"

	"github.com/YarikRevich/HideSeek-Services/pkg/interfaces"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

var (
	instance *BGSaveService
)

type BGSaveService struct {
	*websocket.FastHTTPUpgrader
}

func (b *BGSaveService) Handler(ctx *fasthttp.RequestCtx) {
	b.Upgrade(ctx, func(c *websocket.Conn) {
		defer func() {
			if err := c.Close(); err != nil {
				log.Fatalln(err)
			}
		}()
		for {
			_, _, err := c.ReadMessage()
			if err != nil{
				log.Fatalln(err)
			}
		}
	})
}

func (b *BGSaveService) Send(msg []byte){

}

func UseService() interfaces.Service {
	if instance == nil {
		return new(BGSaveService)
	}
	return instance
}
