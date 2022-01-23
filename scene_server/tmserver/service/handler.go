package service

import (
	"context"

	"gorpc/common/rpc"
	"gorpc/common/types"
	"gorpc/scene_server/tmserver/core"
)

func (s *CoreService) DBOperation(input rpc.Request) (interface{}, error) {

	var ctx core.ContextParams

	reply := types.OPReply{}
	err := input.Decode(&ctx)
	if nil != err {
		reply.Message = err.Error()
		return &reply, nil
	}
	ctx.Context = context.Background()
	ctx.ListenIP = s.listenIP

	return s.core.ExecuteCommand(ctx, input)

}

func (s *CoreService) WatchTransaction(input rpc.Request, stream rpc.ServerStream) (err error) {
	ch := make(chan *types.Transaction, 100)
	s.core.Subscribe(ch)
	defer s.core.UnSubscribe(ch)
	for txn := range ch {
		if err = stream.Send(txn); err != nil {
			return err
		}
	}
	return nil
}
