package main

import (
	"fmt"
	"sync"
)

type ContadorSeguro struct {
	mu sync.Mutex
	v  int
}
func main() {
	var wg = new(sync.WaitGroup)
	c := ContadorSeguro{}

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go c.Inc(wg)
	}
	wg.Wait()
	fmt.Println(c.Get())
}

func (c *ContadorSeguro) Inc(wg *sync.WaitGroup) {
	defer wg.Done()

	c.mu.Lock()
	c.v = c.v + 1  // read + write
	c.mu.Unlock()
}

func (c *ContadorSeguro) Get() int {
	defer c.mu.Unlock()

	c.mu.Lock()
	return c.v
}

