package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const SampleSize = 100
const NumberOfInvocations = 100000

var x int

func mutexExp() {
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < NumberOfInvocations; i++ {
		wg.Add(1)
		go func(mx *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			mx.Lock()
			x++
			mx.Unlock()
		}(&mx, &wg)
		wg.Add(1)
		go func(mx *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			mx.Lock()
			x--
			mx.Unlock()
		}(&mx, &wg)
		wg.Wait()
	}
}
func chanExpInt() {
	ch := make(chan int, 1)
	wg := sync.WaitGroup{}

	for i := 0; i < NumberOfInvocations; i++ {
		wg.Add(1)
		go func(ch chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- 1
			x++
			<-ch
		}(ch, &wg)
		wg.Add(1)
		go func(ch chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- 1
			x--
			<-ch
		}(ch, &wg)
		wg.Wait()
	}
}
func chanExpBool() {
	ch := make(chan bool, 1)
	wg := sync.WaitGroup{}

	for i := 0; i < NumberOfInvocations; i++ {
		wg.Add(1)
		go func(ch chan bool, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- true
			x++
			<-ch
		}(ch, &wg)
		wg.Add(1)
		go func(ch chan bool, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- true
			x--
			<-ch
		}(ch, &wg)
		wg.Wait()
	}
}
func chanExpStruct() {
	ch := make(chan struct{}, 1)
	wg := sync.WaitGroup{}

	for i := 0; i < NumberOfInvocations; i++ {
		wg.Add(1)
		go func(ch chan struct{}, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- struct{}{}
			x++
			<-ch
		}(ch, &wg)
		wg.Add(1)
		go func(ch chan struct{}, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- struct{}{}
			x--
			<-ch
		}(ch, &wg)
		wg.Wait()
	}
}

func main() {

	runtime.GOMAXPROCS(1)

	for i := 0; i < SampleSize; i++ {
		startTime := time.Now()
		//chanExpBool()
		//chanExpInt()
		//chanExpStruct()
		mutexExp()
		endTime := time.Now()
		fmt.Println(endTime.Sub(startTime).Milliseconds())
	}
}

//mutexExp()
//chanExpBool()
