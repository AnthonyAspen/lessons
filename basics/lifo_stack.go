package main

import "fmt"


func main(){
  var array []string
  array = append(array,"first")
  array = append(array,"second")

  for len(array)>0{
    n := len(array) - 1
    fmt.Println(array[n])
    array = array[:n]
  }
}
