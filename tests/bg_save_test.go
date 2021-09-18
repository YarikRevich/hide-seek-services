package bgsave_test

import (
	"testing"
	"time"

	"github.com/YarikRevich/HideSeek-Services/internal/server"
	bgsave "github.com/YarikRevich/HideSeek-Services/pkg/bg_save"
	"github.com/franela/goblin"
)

var (
	g *goblin.G
)

func init(){
	server.Init()
	go server.Run()

	time.Sleep(2 * time.Second)
}

func TestBgSaveSend(t *testing.T) {
	g = goblin.Goblin(t)

	g.Describe("...", func() {

		g.It("test", func() {
			bgsave.UseService().Send([]byte(""))
		})
	})
}

func TestBgSaveHandler(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("...", func() {

		g.It("test", func() {
			bgsave.UseService().Send([]byte(""))
		})
	})
}
