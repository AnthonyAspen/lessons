package main

import "fmt"

func sum(arr []int,c chan int){

  var sum int

  for _,v := range arr {
    sum += v
  }
  c <- sum
  c <- sum + sum
  c <- sum + sum + sum
  fmt.Println("hello")
}

func main(){
  c := make(chan int)
  arr := []int{1,2,3,4,5,6,7}

  go sum(arr,c)
  answer1 := <-c
  answer2 := <-c
  answer3 := <-c
  fmt.Println(answer1,answer2,answer3)


}
