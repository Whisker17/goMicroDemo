package main

import (
	"context"
	"fmt"
	"github.com/Whisker17/goMicroDemo/proto/model"
	"github.com/Whisker17/goMicroDemo/proto/rpcapi"
	"github.com/Whisker17/goMicroDemo/util"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"os"
	"os/signal"
)

func main() {
	// 初始化服务
	service := micro.NewService(
		//micro.Registry(reg),
	)

	service.Init()
	service.Client().Init(client.Retries(3),
		client.PoolSize(5))
	sayClent := rpcapi.NewSayService(util.ServiceName, service.Client())

	SayHello(sayClent)
	NotifyTopic(service)

	st := make(chan os.Signal)
	signal.Notify(st, os.Interrupt)

	<-st
	fmt.Println("server stopped.....")
}

func SayHello(client rpcapi.SayService) {
	rsp, err := client.Hello(context.Background(), &model.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
}

func NotifyTopic(service micro.Service) {
	p := micro.NewPublisher(util.Topic, service.Client())
	p.Publish(context.TODO(), &model.SayParam{Msg: util.RandomStr(util.Random(3, 10))})
}