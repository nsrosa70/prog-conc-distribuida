package main

import (
	"fmt"
   "multithread/synccond/mutexqueue"
   "sync"
)

func main() {

   var wg = new(sync.WaitGroup)

   //var q = condqueue.NewCondQueue(3)
   var q = mutexqueue.NewMutexQueue(3)

   wg.Add(5)

   go q.Push(1,wg)
   go q.Push(2,wg)
   go q.Push(3,wg)
   go q.Push(4,wg)
   go q.Pop(wg)

   fmt.Println(q)
}
