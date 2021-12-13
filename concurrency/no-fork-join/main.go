package main

import (
	"fmt"
	"time"
)



func main(){
  go work()
  time.Sleep(time.Millisecond)
  fmt.Println("done waiting...Quiting")

}

func work(){
  time.Sleep(500*time.Millisecond)
  fmt.Println("print out some stuff")
}
