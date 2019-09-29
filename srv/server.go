package main

import (
	"github.com/Whisker17/goMicroDemo/handler"
	"github.com/Whisker17/goMicroDemo/proto/rpcapi"
	"github.com/Whisker17/goMicroDemo/util"
	"github.com/micro/go-micro"
	"time"
)

func main() {
	// 初始化服务
	service := micro.NewService(
		micro.Name(util.ServiceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20),

		//micro.Registry(reg),
	)

	service.Init()
	// 注册 Handler
	rpcapi.RegisterSayHandler(service.Server(), new(handler.Say))

	// Register Subscribers
	//if err := server.Subscribe(server.NewSubscriber(util.Topic, subscriber.Handler)); err != nil {
	//	panic(err)
	//}

	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
