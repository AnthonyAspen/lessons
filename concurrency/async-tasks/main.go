package main

import (
	"fmt"
	"time"
)

func main(){
  now := time.Now()
  done := make(chan struct{})
  go  task0(done);
  go  task1(done);
  go  task2(done);
  go  task3(done);
  <-done
  <-done
  <-done
  <-done
  time.Sleep(time.Second)
  fmt.Println("elapsed:",time.Since(now))

}
func task0(done chan struct{}){
  time.Sleep(100*time.Millisecond)
  fmt.Println("task0")

}
func task1(done chan struct{}){
  time.Sleep(200*time.Millisecond)
  fmt.Println("task1")

}
func task2(done chan struct{}){
  fmt.Println("task2")

}
func task3(done chan struct{}){
  time.Sleep(100*time.Millisecond)
  fmt.Println("task3")

}
