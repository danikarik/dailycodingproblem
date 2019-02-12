package problem14

import (
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	rnd   = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	min   = -1.0
	max   = 1.0
	mutex = sync.Mutex{}
)

func generate() (float64, float64) {
	mutex.Lock()
	defer mutex.Unlock()
	x := min + rnd.Float64()*(max-min)
	y := min + rnd.Float64()*(max-min)
	return x, y
}

func inCircle(x, y float64) bool {
	return x*x+y*y < 1
}

func estimate(n int) float64 {
	var inside uint64
	insideC := make(chan struct{})
	exitC := make(chan struct{})
	for i := 0; i < n; i++ {
		go func(last bool) {
			if inCircle(generate()) {
				insideC <- struct{}{}
			}
			if last {
				exitC <- struct{}{}
			}
		}(i == n-1)
	}
	for {
		select {
		case <-insideC:
			atomic.AddUint64(&inside, 1)
		case <-exitC:
			close(insideC)
			close(exitC)
			ratio := float64(atomic.LoadUint64(&inside)) / float64(n)
			return math.Round(ratio*4*1000) / 1000
		}
	}
}
