package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func New() Counter {
	return Counter{
		val: atomic.Uint32{},
	}
}

type Counter struct {
	val atomic.Uint32
}

func (c *Counter) Increment() {
	c.val.Store(c.val.Load() + 1)
}

func (c *Counter) Value() uint32 {
	return c.val.Load()
}

func main() {
	c := New()
	for i := 0; i < 100; i++ {
		go func() {
			c.Increment()
		}()
	}

	time.Sleep(time.Second * 2)

	fmt.Println(c.Value())
}
