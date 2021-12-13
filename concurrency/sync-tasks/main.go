package main

import (
	"fmt"
	"time"
)

func main(){
  now := time.Now()
  task0();
  task1();
  task2();
  task3();
  fmt.Println("elapsed:",time.Since(now))

}

func task0(){
  time.Sleep(100*time.Millisecond)
  fmt.Println("task0")

}
func task1(){
  time.Sleep(200*time.Millisecond)
  fmt.Println("task1")

}
func task2(){
  fmt.Println("task2")

}
func task3(){
  time.Sleep(100*time.Millisecond)
  fmt.Println("task3")

}
