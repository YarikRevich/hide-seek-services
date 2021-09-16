package bgsave_test

import (
	"testing"

	bgsave "github.com/YarikRevich/HideSeek-Services/pkg/bg_save"
	"github.com/franela/goblin"
	"github.com/golang/mock"
)

func TestBgSave(t *testing.T){
	g := goblin.Goblin(t)

	g.Describe("...", func(){
		g.It("...", func(){
			bgsave.UseService().Send(...)
		})
	})
}