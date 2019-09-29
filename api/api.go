package main

import (
	"context"
	"encoding/json"
	"github.com/Whisker17/goMicroDemo/proto/model"
	"github.com/Whisker17/goMicroDemo/proto/rpcapi"
	"github.com/Whisker17/goMicroDemo/util"
	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"log"
	"strings"
)

type Say struct {
	Client rpcapi.SayService
}

func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.Hello API request")

	name, ok := req.Get["msg"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest(util.ApiName, "Name cannot be blank")
	}

	response, err := s.Client.Hello(ctx, &model.SayParam{
		Msg: strings.Join(name.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)

	return nil
}

func (s *Say) MyName(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.MyName API request")

	response, err := s.Client.MyName(ctx, &model.SayParam{})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)

	return nil
}

func main() {
	// new一个微服务出来 资源类型设置为api
	service := micro.NewService(
		micro.Name(util.ApiName),
	)

	// 可选 解析命令行
	service.Init()

	// 注册handler
	service.Server().Handle(
		service.Server().NewHandler(
			&Say{Client: rpcapi.NewSayService(util.ServiceName, service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

