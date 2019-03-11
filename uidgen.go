package uidgen

import (
	"sync/atomic"
	"time"
)

type uidgeneratorTs struct {
	ServerID uint64
	Counter  uint64
	Ts       uint64
}

func (g *uidgeneratorTs) GenUID() uint64 {
	newid := atomic.AddUint64(&g.Counter, 1) & 0x00ffffff
	newid |= g.Ts
	newid |= g.ServerID
	return newid
}

func CreateGenerator(serverID uint64) UIDGenerator {
	gen := &uidgeneratorTs{
		ServerID: serverID << 24,
	}
	atomic.StoreUint64(&gen.Counter, 0)

	UpdateTs := func() {
		ts := (uint64(time.Now().UTC().Unix()) & (0xffffffff)) << 32
		gen.Ts = ts
	}

	UpdateTs()
	go func() {
		for {
			time.Sleep(time.Millisecond * 250)
			UpdateTs()
		}
	}()

	return gen
}
