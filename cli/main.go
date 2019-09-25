package main

import (
	"context"
	"fmt"
	"github.com/Whisker17/goMicroDemo/proto/model"
	"github.com/Whisker17/goMicroDemo/proto/rpcapi"
	"github.com/Whisker17/goMicroDemo/util"
	"github.com/lpxxn/gomicrorpc/example2/lib"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"io"
	"os"
	"os/signal"
)

func main() {
	// 我这里用的etcd 做为服务发现
	//reg := etcdv3.NewRegistry(func(op *registry.Options) {
	//	op.Addrs = []string{
	//		"http://192.168.3.156:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379",
	//	}
	//})

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
	GetStreamValues(sayClent)
	TsBidirectionalStream(sayClent)

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

// test stream
func GetStreamValues(client rpcapi.SayService) {
	rspStream, err := client.Stream(context.Background(), &model.SRequest{Count: 10})
	if err != nil {
		panic(err)
	}

	idx := 1
	for {
		rsp, err := rspStream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		fmt.Printf("test stream get idx %d  data  %v\n", idx, rsp)
		idx++
	}
	// close the stream
	if err := rspStream.Close(); err != nil {
		fmt.Println("stream close err:", err)
	}
	fmt.Println("Read Value End")
}

func TsBidirectionalStream(client rpcapi.SayService) {
	rspStream, err := client.BidirectionalStream(context.Background())
	if err != nil {
		panic(err)
	}
	for i := int64(0); i < 7; i++ {
		if err := rspStream.Send(&model.SRequest{Count: i}); err != nil {
			fmt.Println("send error", err)
			break
		}
		rsp, err := rspStream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		fmt.Printf("test stream get idx %d  data  %v\n", i, rsp)
	}
	// close the stream
	if err := rspStream.Close(); err != nil {
		fmt.Println("stream close err:", err)
	}
	fmt.Println("TsBidirectionalStream: Read Value End")
}

func NotifyTopic(service micro.Service) {
	p := micro.NewPublisher(util.Topic, service.Client())
	p.Publish(context.TODO(), &model.SayParam{Msg: lib.RandomStr(lib.Random(3, 10))})
}
