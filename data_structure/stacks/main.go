package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Stack struct{
  items []int
}

// Push will add a balue at the top
func (s *Stack) Push(i int){
  s.items = append(s.items,i)
}
// Pop will remove a value at the top
func (s *Stack) Pop() (value int,err error){
  // if the stack is empty and a person tries to pop
  if len(s.items) == 0{
    return -1,errors.New("The stack is empty")
  }
  l := len(s.items)-1

  value = s.items[l]
  s.items = s.items[:l]
  return value,nil
}

func main(){
  stack := Stack{}
  for i := 0;i<10;i++{
    stack.Push(rand.Intn(100-1)+1)
  }
  fmt.Println(stack)
  fmt.Println(stack.Pop())
}
