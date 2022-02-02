package main

import "fmt"


func main(){

  slice := make([]int,2,2)
  slice = append(slice,0)
  slice = append(slice,1)
  fmt.Printf("pointer %p cap:%d\n",&slice,cap(slice))
  slice = append(slice,2)
  fmt.Printf("pointer %p cap:%d\n",&slice,cap(slice))



  sliceAnother := []int{0,1,2,3}

  fmt.Printf("result %d another result with : %d \n",test(sliceAnother...),test(sliceAnother[:]...))

  slice = append(slice,sliceAnother...)
  fmt.Println(slice)
  remove(slice,0)
  fmt.Println(slice)
  remove(slice,0)
  fmt.Println(slice)
  remove(slice,0)
  fmt.Println(slice)



}
// remove an element from a slice 
func remove (slice []int,s int) []int{
  return append(slice[:s],slice[s+1:]...)
}
// testing how append(slice []type, elems ... Type) []type works

func test(args ... int) (result int){

  for arg := range args {
    result += arg
  }

 return  
}
