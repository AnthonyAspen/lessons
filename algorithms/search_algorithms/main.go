package main

import "fmt"


/*
 Linear Search
 This technique pass over the list of elements, by using the index to move from
 the beginning of the list to the end. Each element is examined and if it does
 not match the search item, the next item is examined. By hopping from one item
 to its next, the list is passed over sequentially.
*/
func linearSearch(datalist []int,key int) bool{
  for _,v := range datalist{
    if v == key {
      return true 
    }
  }
  return false

}
func binarySearch(datalist []int,key int) bool{
  low,high := 0, len(datalist)/2

  for low <= high {
    median := (low + high)/2
    
    if datalist[median] < key {
      low = median + 1
    }else{
      high = median - 1
    }
  }

  if low == len(datalist) || datalist[low] != key {
    return false
  }

  return true 
}

func main(){
  // linear search
  randItems := []int{40,30,60,10,70,50,20}
  sortItems := []int{10,20,30,40,50,60,70}
  
  fmt.Printf("Linear search of randItems did find his element: %v\n",linearSearch(randItems,50))
  fmt.Printf("Binary search of sortItems did find his element: %v\n",binarySearch(sortItems,50))
}
