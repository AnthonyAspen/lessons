package main

import (
	"errors"
	"fmt"
	"log"
)

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
  array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int){
  h.array = append (h.array,key)
  h.maxHeapifyUp(len(h.array) - 1)
}
 
// Extract returns the largest key, and removes it from the heap
// TODO extract doesn't work properly
func (h *MaxHeap) Extract() (int,error){
  extracted := h.array[0]
  l := len(h.array) - 1
  if l == -1 {
    var err = errors.New("Cannot extract because the array length is 0")
    return 0, err
  }
  h.array[0] = h.array[l]
  h.array = h.array[:l]
  return extracted,nil
}
// maxHeapifyUp heapifies frrom bottom to top
// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) maxHeapifyUp(index int){
  // when the parent is smaller than the current index
  for h.array[parent(index)] < h.array[index]{
    // then we're going to swap
    h.swap(parent(index),index)
    // after we swapped we're going to perform the same loop on the parent
    index = parent(index)


  }
}
// maxHeapifyDown
func (h *MaxHeap) maxHeapifyDown(index int){
  lastIndex := len(h.array) -1 
  l,r := left(index),right(index)
  childToCompare := 0
  
  // loop while index has at least one child
  for l <= lastIndex{

    if l == lastIndex{
    // when left child is the only child
      childToCompare = l

    }else if h.array[l] > h.array[r]{
    // when left child is larger
    childToCompare = l

  }else {
    //when right child is larger
    childToCompare = r
  }

  // compare array value of current index to larger child and swap if smaller
  if h.array[index] < h.array[childToCompare]{
    h.swap(index,childToCompare)
    index = childToCompare
    l,r = left(index), right(index)
  }else{
    return
  }

  }
}
// swap keys in the array
func (h *MaxHeap) swap(i1,i2 int){
  h.array[i1],h.array[i2] = h.array[i2],h.array[i1]
}

// get the parent index
func parent(i int) int{
  return (i-1)/2
}

// get the left child index
func left(i int) int{
  return 2*i + 1
}

// get the right child index
func right(i int) int{
  return 2*i + 2
}

// main function
func main(){
  m := &MaxHeap{}
  buildHeap := []int{10,20,30,40,34,12,9,3,5}

  for _,v := range buildHeap{
    m.Insert(v)
    fmt.Println(m)
  }
  for i := 0; i < 5; i ++{
    _,err := m.Extract()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(m)
  }
}



