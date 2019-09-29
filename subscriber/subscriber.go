package subscriber

import (
	"context"
	"fmt"
	"github.com/Whisker17/goMicroDemo/proto/model"
)

func Handler(ctx context.Context, msg *model.SayParam) error {
	fmt.Printf("Received message: %s \n", msg.Msg)
	return nil
}
