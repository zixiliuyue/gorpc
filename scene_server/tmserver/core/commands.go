package core

import (
	"gorpc/common/rpc"
	"gorpc/common/types"
)

var (
	// GCommands global db operation command map definition
	GCommands = &commands{cmds: nil}
)

// Command db operation definition
type Command interface {
	Execute(ctx ContextParams, decoder rpc.Request) (*types.OPReply, error)
}

type commands struct {
	cmds map[types.OPCode]Command
}

func (c *commands) SetCommand(opCode types.OPCode, cmd Command) {
	if nil == c.cmds {
		c.cmds = make(map[types.OPCode]Command)
	}

	c.cmds[opCode] = cmd
}
