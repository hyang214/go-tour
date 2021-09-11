package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter interface {
	Incr(s string)
	Value(s string) int
}

type UnsafeCounter struct {
	v map[string]int
}

func (c *UnsafeCounter) Incr(s string)  {
	c.v[s] ++
}

func (c *UnsafeCounter) Value(s string) int  {
	return c.v[s]
}

type SafeCounter struct {
	m sync.Mutex
	v map[string]int
}

func (c *SafeCounter) Incr(s string)  {
	c.m.Lock()
	c.v[s] ++
	c.m.Unlock()
}

func (c *SafeCounter) Value(s string) int  {
	c.m.Lock()
	defer c.m.Unlock()
	return c.v[s]
}

func main() {
	// lead to "fatal error: concurrent map writes"
	unsafe := &UnsafeCounter{v: make(map[string]int)}
	Test(unsafe)

	safe := &SafeCounter{v: make(map[string]int)}
	Test(safe)
}

func Test(c Counter)  {
	for i := 0; i < 100; i++ {
		go c.Incr("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
