package daemon

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func RunDaemon(){
	if _, ok := os.LookupEnv("DAEMONIZED"); !ok {
		c := exec.Command(os.Args[0])
		c.Env = append(c.Env, "DAEMONIZED")
		if err := c.Start(); err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}
	syscall.Umask(022)
	workPath, err := filepath.Abs(os.Args[0])
	if err != nil{
		log.Fatalln(err)
	}
	if err := os.Chdir(workPath); err != nil {
		log.Fatalln(err)
	}
	if _, err := syscall.Setsid(); err != nil {
		log.Fatalln(err)
	}
}