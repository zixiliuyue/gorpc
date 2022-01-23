package core

import (
	"fmt"

	"gorpc/common/blog"
	"gorpc/common/rpc"
	"gorpc/common/types"
)

// Core core operation methods
type Core interface {
	ExecuteCommand(ctx ContextParams, input rpc.Request) (*types.OPReply, error)
	Subscribe(chan *types.Transaction)
	UnSubscribe(chan<- *types.Transaction)
}

type core struct {
	enable bool
}

// New create a core instance
func New() Core {
	// for _, cmd := range GCommands.cmds {
	// 	switch tmp := cmd.(type) {
	// 	case SetTransaction:
	// 		tmp.SetTxn(txnMgr)
	// 	case SetDBProxy:
	// 		tmp.SetDBProxy(db)
	// 	}
	// }
	return &core{}
}

func (c *core) ExecuteCommand(ctx ContextParams, input rpc.Request) (*types.OPReply, error) {

	blog.V(5).Infof("RDB operate. info:%#v", ctx.Header)

	cmd, ok := GCommands.cmds[ctx.Header.OPCode]
	if !ok {
		blog.ErrorJSON("RDB operate unkonwn operation")
		reply := types.OPReply{}
		reply.Message = fmt.Sprintf("unknown operation, invalid code: %d", ctx.Header.OPCode)
		return &reply, nil
	}

	// session := c.txn.GetSession(ctx.Header.TxnID)
	// if nil == session {
	// 	reply := &types.OPReply{}
	// 	reply.Message = "session not found"
	// 	return reply, nil
	// }
	// ctx.Session = session.Session

	reply, err := cmd.Execute(ctx, input)
	if err != nil {
		blog.Errorf("[MONGO OPERATION] failed: %v, cmd: %s", err, input)
	}
	return reply, err

}

func (c *core) Subscribe(ch chan *types.Transaction) {
	// c.txn.Subscribe(ch)
}

func (c *core) UnSubscribe(ch chan<- *types.Transaction) {
	// c.txn.UnSubscribe(ch)
}
