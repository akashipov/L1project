package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Counter struct {
	v *atomic.Int64
}

func NewCounter(value int64) *Counter {
	c := atomic.Int64{}
	c.Store(int64(value))
	return &Counter{&c}
}

func (c *Counter) Count() {
	c.v.Add(1)
}

func main() {
	c := NewCounter(0)
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(c.v.Load())
		}
	}()
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			c.Count()
		}
	}()
	time.Sleep(time.Second * 5)

}
