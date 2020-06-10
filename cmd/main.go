package main

import (
	"fmt"
	"github.com/pangxieke/sendmail/config"
	"github.com/pangxieke/sendmail/log"
	"github.com/pangxieke/sendmail/model"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}
	if err := log.Init(config.Server.LogFile); err != nil {
		panic(err)
	}
	if err := model.Init(); err != nil {
		panic(err)
	}

	fmt.Println("server mail start")

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go model.SyncRedis()

	fmt.Println("exit", <-errc)
}
