package main

import (
	"fmt"
	"monitors/wordmonitor/impl"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	m := &impl.Words{}
	m.Init()

	wg.Add(2)
	go func() {
		defer wg.Done()
		m.SetData("Angad")
	}()

	go func() {
		defer wg.Done()
		m.SetData("Sharma")
	}()
	wg.Wait()

	fmt.Println(m.GetData())
}
