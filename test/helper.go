package test

import (
	"github.com/pangxieke/sendmail/config"
	"github.com/pangxieke/sendmail/log"
	"github.com/pangxieke/sendmail/model"
)

func init() {
	err := config.Init("../test")
	if err != nil {
		panic(err)
	}
	if err := log.Init(config.Server.LogFile); err != nil {
		panic(err)
	}
	if err := model.Init(); err != nil {
		panic(err)
	}
	return
}
