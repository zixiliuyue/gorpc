package command

import (
	"gorpc/common/blog"
	// "gorpc/src/storage/mongodb"
	"gorpc/common/rpc"
	"gorpc/common/types"
	"gorpc/scene_server/tmserver/core"
)

func init() {
	core.GCommands.SetCommand(types.OPCountCode, &count{})
}

type count struct {
}

func (d *count) Execute(ctx core.ContextParams, decoder rpc.Request) (*types.OPReply, error) {

	msg := types.OPDeleteOperation{}
	reply := &types.OPReply{}
	reply.RequestID = ctx.Header.RequestID
	if err := decoder.Decode(&msg); nil != err {
		reply.Message = err.Error()
		return reply, err
	}
	blog.V(4).Infof("[MONGO OPERATION] %+v", &msg)
	reply.Count = 1111111
	reply.Success = true
	return reply, nil
}
