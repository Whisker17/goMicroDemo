package handler

import (
	"context"
	"fmt"
	"github.com/Whisker17/goMicroDemo/proto/model"
	"github.com/Whisker17/goMicroDemo/proto/rpcapi"
)


type Say struct{}

var _ rpcapi.SayHandler = (*Say)(nil)

func (s *Say) Hello(ctx context.Context, req *model.SayParam, rsp *model.SayResponse) error {
	fmt.Println("received", req.Msg)
	rsp.Header = make(map[string]*model.Pair)
	rsp.Header["name"] = &model.Pair{Key: 1, Values: "abc"}

	rsp.Msg = "hello " + req.Msg
	rsp.Values = append(rsp.Values, "a", "b")
	rsp.Type = model.RespType_DESCEND

	return nil
}

func (s *Say) MyName(ctx context.Context, req *model.SayParam, rsp *model.SayResponse) error {
	fmt.Println("received", req.Msg)
	rsp.Header = make(map[string]*model.Pair)
	rsp.Header["name"] = &model.Pair{Key: 2, Values: "def"}
	rsp.Msg = "whisker"
	rsp.Values = append(rsp.Values, "d", "e")
	rsp.Type = model.RespType_DESCEND
	return nil
}
