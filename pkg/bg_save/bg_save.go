package bgsave

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/YarikRevich/HideSeek-Services/pkg/common"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

var (
	instance *BGSaveService
)

type BGSaveService struct {
	serviceName string
	url         string

	websocket.FastHTTPUpgrader
}

func (b *BGSaveService) Handler() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		err := b.Upgrade(ctx, func(c *websocket.Conn) {
			defer func() {
				if err := c.Close(); err != nil {
					log.Fatalln(err)
				}
			}()
			for {
				_, _, err := c.ReadMessage()
				
				if websocket.IsCloseError(err, websocket.CloseNoStatusReceived) {
					break
				}

				if err != nil {
					log.Fatalln(err)
				}

			}
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (b *BGSaveService) Send(msg []byte) {
	c, _, err := websocket.DefaultDialer.Dial(b.url, nil)
	if err != nil {
		log.Fatalln()
	}

	var buff bytes.Buffer
	if _, err = zlib.NewWriter(&buff).Write(msg); err != nil {
		log.Fatalln(err)
	}
	if err = c.WriteMessage(websocket.BinaryMessage, buff.Bytes()); err != nil {
		log.Fatalln(err)
	}
}

func (b *BGSaveService) SendAsync(msg []byte) {
	go b.Send(msg)
}

func (b *BGSaveService) SetURL(host string) {
	b.url = fmt.Sprintf("ws://%s/%s", host, b.serviceName)
}

func (b *BGSaveService) Name() string {
	return b.serviceName
}

func UseService() common.Service {
	if instance == nil {
		instance = &BGSaveService{
			serviceName: strings.ToLower(reflect.TypeOf(new(BGSaveService)).Elem().Name()),
		}
	}
	return instance
}
