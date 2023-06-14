package main

import (
	"fmt"
	"sync"
	"time"
)

const SampleSize = 30
const NumberOfInvocations = 100000

var v int

func MutexExp() {
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < NumberOfInvocations; i++ {
		wg.Add(1)
		go func(mx *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			mx.Lock()
			v++
			mx.Unlock()
		}(&mx, &wg)

		wg.Add(1)
		go func(mx *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			mx.Lock()
			v--
			mx.Unlock()
		}(&mx, &wg)
		wg.Wait()
	}
}

func ChanExp() {
	//ch := make(chan bool, 1)
	ch := make(chan string, 1)
	wg := sync.WaitGroup{}
	//ch <- true
	ch <- "ok"

	for i := 0; i < NumberOfInvocations; i++ {
		wg.Add(1)
		//go func(ch chan bool, wg *sync.WaitGroup) {
		go func(ch chan string, wg *sync.WaitGroup) {
			defer wg.Done()
			<-ch
			v++
			//ch <- true
			ch <- "ok"
		}(ch, &wg)

		wg.Add(1)
		//go func(mx chan bool, wg *sync.WaitGroup) {
		go func(mx chan string, wg *sync.WaitGroup) {
			defer wg.Done()
			<-ch
			v--
			//ch <- true
			ch <- "ok"
		}(ch, &wg)
		wg.Wait()
	}
}

func main() {

	for n := 0; n < SampleSize; n++ {
		start := time.Now()
		//MutexExp()
		ChanExp()
		end := time.Now()
		fmt.Println(end.Sub(start).Milliseconds())
	}
}
