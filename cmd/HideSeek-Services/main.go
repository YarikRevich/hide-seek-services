package main

import (
	"flag"

	"github.com/YarikRevich/HideSeek-Services/internal/daemon"
	"github.com/YarikRevich/HideSeek-Services/internal/server"
	"github.com/YarikRevich/HideSeek-Services/tools/cli"
)

func init() {
	flag.Parse()

	server.Init()
}

func main() {
	if cli.IsDaemon() {
		daemon.RunDaemon()
	} else {
		server.Wait(server.Run())
	}
}
