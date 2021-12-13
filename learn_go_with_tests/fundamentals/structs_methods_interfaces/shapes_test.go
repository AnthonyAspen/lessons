package main

import "testing"

type Rectangle struct {
  Width float64
  Height float64
}

func Perimeter(rectangle Rectangle) float64{
  return 2 * (rectangle.Height + rectangle.Width)
}

func Area(rectangle Rectangle) float64{
  return rectangle.Height + rectangle.Width
}
func TestPerimeter(t *testing.T){
  rectangle := Rectangle{10.0,10.0}
  got := Perimeter(rectangle)
  want := 40.0
  if got != want {
    t.Errorf("got %.2f want %.2f",got,want)
  }
}
func TestArea(t *testing.T){
  rectangle := Rectangle{12.0,6.0}
  got := Area(rectangle)
  want := 72.0

  if got != want {
    t.Errorf("got %.2f want %.2f",got,want)
  }
}
