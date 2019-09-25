package main

import (
	"context"
	"fmt"
	"github.com/Whisker17/goMicroDemo/proto"
	"github.com/micro/go-micro"
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
		//micro.Registry(reg),
		micro.Name("whisker.srv.eg1"),
	)

	// 2019年源码有变动默认使用的是mdns面不是consul了
	// 如果你用的是默认的注册方式把上面的注释掉用下面的
	/*
		// 初始化服务
		service := micro.NewService(
			micro.Name("lp.srv.eg1"),
		)
	*/
	service.Init()

	sayClent := model.NewSayService("whisker.srv.eg1", service.Client())

	rsp, err := sayClent.Hello(context.Background(), &model.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)

}
