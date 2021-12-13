package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randNumsGenerator(n int) <-chan int{
  r := rand.New(rand.NewSource(time.Now().UnixNano()))

  out := make(chan int)
  
  go func(){
    for i := 0; i < n; i++ {
      out <- r.Intn(n)
    }
    close(out)
  }()
  return out
}

func main(){
  for num := range randNumsGenerator(10){
    fmt.Println(num)
  }

}

