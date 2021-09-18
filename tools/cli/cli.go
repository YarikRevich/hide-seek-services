package cli

import "flag"

var (
	daemon = flag.Bool("daemon", false, "Daemon mode")
)

func IsDaemon() bool{
	return *daemon
}