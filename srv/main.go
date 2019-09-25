package main

import (
	"github.com/Whisker17/goMicroDemo/handler"
	"github.com/Whisker17/goMicroDemo/proto/rpcapi"
	"github.com/Whisker17/goMicroDemo/subscriber"
	"github.com/Whisker17/goMicroDemo/util"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"time"
)

func main() {
	// 我这里用的etcd 做为服务发现
	//reg := etcdv3.NewRegistry(func(op *registry.Options) {
	//	op.Addrs = []string{
	//		"http://192.168.3.34:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379",
	//	}
	//})

	// 初始化服务
	service := micro.NewService(
		micro.Name(util.ServiceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20),
		//micro.Registry(reg),
		//micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)


	service.Init()
	// 注册 Handler
	rpcapi.RegisterSayHandler(service.Server(), new(handler.Say))

	// Register Subscribers
	if err := server.Subscribe(server.NewSubscriber(util.Topic, subscriber.Handler)); err != nil {
		panic(err)
	}

	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}

func PrometheusBoot() {
	http.Handle("metrics",prometheus.Handler())

	go func() {
		err := http.ListenAndServe("192.168.3.156:8085",nil)
		if err != nil {
			log.Fatal("ListenAndServe:",err)
		}
	}()
}

