package main

import (
	"chat/conf"
	"chat/mq"
	"chat/router"
	"chat/script"
	"chat/service"
	"context"
)

func main() {
	conf.Init()
	mq.InitRabbitMQ()
	loadingScript()
	go service.Manager.Start()
	r := router.NewRouter()
	_ = r.Run(conf.HttpPort)
}
func loadingScript() {
	ctx := context.Background()
	go script.TaskCreateSync(ctx)
}
