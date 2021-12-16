package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// type Tree struct {
// 	Left  *Tree
// 	Value int
// 	Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
  defer close(ch)
  var walk func(t*tree.Tree)
  walk = func(t *tree.Tree) {
    if t == nil{
      return
    }
    walk(t.Left)
    ch <- t.Value
    walk(t.Right)
  }
  walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
  ch1 := make(chan int)
  ch2 := make(chan int)

  firstSlice := make([]int,0,10)
  secondSlice:= make([]int,0,10)

  go Walk(t1,ch1)
  go Walk(t2,ch2)
  for i := range ch1{
    firstSlice = append(firstSlice,i)
  }
  for i := range ch2{
    secondSlice = append(secondSlice,i)
  }
  for _,v := range firstSlice{
    fmt.Printf("cap : %v, len: %v, num : %v\n",cap(firstSlice),len(firstSlice),v)
  }
  return CompareSlices(firstSlice,secondSlice)
}

// func to compare 2 slices 
func CompareSlices(s1,s2 []int) bool{
  if len(s1) != len(s2){
    return false
  } 
  for i := range s1 {
    if s1[i] != s2[i]{
      return false
    }
  }
  return true
}

func main() {
firstTree := tree.New(1)
secondTree := tree.New(2)
fmt.Println(Same(firstTree,secondTree))


  
}

