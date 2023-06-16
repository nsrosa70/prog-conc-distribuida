package main

import "time"

func main() {
   var x[] int

   go func(){
      time.Sleep(100 * time.Millisecond)
      x = make([]int,10)
   }()

   go func(){
      x = make([]int,100)
   }()

   time.Sleep(100 * time.Millisecond)
   x[50] = 10
}
