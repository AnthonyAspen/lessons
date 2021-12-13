package main

import "fmt"


func doSomething(slice []int,y int){
  slice[0] = slice[1]
  y = 10


}

func main(){
  y := 0
  x := []int{1,2}
  doSomething(x,y)
  fmt.Println(x,y)

}
