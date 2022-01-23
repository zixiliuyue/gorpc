package core

import (
	"context"
	"time"

	// "configcenter/src/storage/mongodb"
	"gorpc/common/types"
)

// ContextParams the logic function params
type ContextParams struct {
	context.Context
	// Session  mongodb.Session
	ListenIP string
	Header   types.MsgHeader `bson:"msgheader"`
}

// Deadline overwrite Context Deadline methods
func (c ContextParams) Deadline() (deadline time.Time, ok bool) {
	return c.Context.Deadline()
}

// Done overwrite Context Done methods
func (c ContextParams) Done() <-chan struct{} {
	return c.Context.Done()
}

// Err overwrite Context Err methods
func (c ContextParams) Err() error {
	return c.Context.Err()
}

// Value overwrite Context Value methods
func (c ContextParams) Value(key interface{}) interface{} {
	return c.Context.Value(key)
}
