package main

import "testing"

func SumArray(numbers [5] int) int{
  sum := 0
  for _, number := range numbers{
    sum += number
  }
  return sum

}
func SumSlice(numbers [] int) int{
  sum := 0
  for _, number := range numbers{
    sum += number
  }
  return sum

}
func TestSumArray(t *testing.T){
  numbers := [...]int{1,2,3,4,5}
  got := SumArray(numbers)
  want := 15
  if got != want{
    t.Errorf("got %d want %d given,%v",got,15,numbers)
  }

}
/// Slices
func TestSum(t *testing.T) {
    t.Run("collection of 5 numbers", func(t *testing.T) {
        numbers := [5]int{1, 2, 3, 4, 5}
        got := SumArray(numbers)
        want := 15
        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })
    t.Run("collection of any size", func(t *testing.T) {
        numbers := []int{1, 2, 3}
        got := SumSlice(numbers)
        want := 6
        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })
}
