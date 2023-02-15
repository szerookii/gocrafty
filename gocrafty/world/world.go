package world

import "sync/atomic"

var cEid int32

func NextEID() int32 {
	return atomic.AddInt32(&cEid, 1)
}

type World struct {
}
