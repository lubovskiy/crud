package helpers

import (
	"math/rand"
	"sync"
	"time"
)

type SeedRand struct {
	r   *rand.Rand
	mux sync.Mutex
}

var sr SeedRand

func init() {
	sr = SeedRand{
		r:   rand.New(rand.NewSource(time.Now().UnixNano())),
		mux: sync.Mutex{},
	}
}

func (sr *SeedRand) SafeInt(i int) int {
	sr.mux.Lock()
	res := sr.r.Intn(i)
	sr.mux.Unlock()
	return res
}

func Int(min, max int) int {
	return min + sr.SafeInt(max-min)
}
