package trivial

import "sync"

type Counter struct {
	count int64
	lock  sync.Mutex
}

func (c *Counter) Increment() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.count++
}

func (c *Counter) Count() int64 {
	return c.count
}
