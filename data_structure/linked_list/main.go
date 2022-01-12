package main

import (
	"errors"
	"fmt"
)

type node struct {
  data int
  next *node
}

type LinkedList struct {
  head *node
  length int
}

func (linkedList *LinkedList) Prepend(a ...*node){
  for _,v := range a{
    second := linkedList.head
    linkedList.head = v
    linkedList.head.next = second
    linkedList.length++
  }
}
// without * because I don't want to change the linked list, just display it
func(linkedList LinkedList) Display(){
  for linkedList.head != nil{
    fmt.Printf("%v -> ",linkedList.head.data)
    linkedList.head = linkedList.head.next
  }
  fmt.Println()
}
func(linkedList *LinkedList) Delete(value int) (error){
  // if the list is empty
  if linkedList.length == 0{
    return  errors.New("The linked list is empty")
  }
  // if we're trying to delete the head of the linked list
  if linkedList.head.data == value{
    linkedList.head = linkedList.head.next
    linkedList.length--
    return nil
  }
  prevToDelete := linkedList.head
  for prevToDelete.next.data != value{
    // if the element does not exist
    if prevToDelete.next.next == nil{
      return errors.New("The element not found ")
    }
    prevToDelete = prevToDelete.next
  }
  prevToDelete.next = prevToDelete.next.next
  linkedList.length--
  

  return nil
}
func main(){
  linkedList := LinkedList{}
  for i := 10;i<100;i= i+i{
    tempNode := &node{data:i}
    linkedList.Prepend(tempNode)

  }
  linkedList.Display()
  linkedList.Delete(34)
  linkedList.Display()
  linkedList.Delete(344)
  linkedList.Display()

}
