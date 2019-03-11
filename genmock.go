package uidgen

import (
	"sync"
)

// Mock generator. Count starts from one.
type MockGen struct {
	Counter uint64
	mu      sync.Mutex
}

func (g *MockGen) GenUID() uint64 {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Counter += 1
	newid := g.Counter
	return newid
}
