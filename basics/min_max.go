package main

import (
  "fmt"
)

func max(x int,y int) int{
  if x > y{
    return x
  }
  return y
}
func min(x int,y int) int{
  if x > y{
    return y
  }
  return x
}
func main(){
  x,y := 15,20
  y,x = x,y
  fmt.Println(max(x,y))
  fmt.Println(min(x,y))

}
