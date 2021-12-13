package main

import "fmt"

func reverse(slice []int) (reversed []int){
  for len(slice) > 0{
    n := len(slice) - 1
    reversed = append(reversed,slice[n])
    slice = slice[:n]
  }

   return 
}
// func reverse_without_var(slice []int) []int{
//   for a,b :=0, len(slice) -1; a < b; a,b = a+1,b-1{

//   }
// }
func main(){
  data := []int{5,4,3,2,1}
  data = reverse(data)
  fmt.Println(data)

}
