package main

import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx,dy int) [][]uint8{
  fmt.Printf("dx = %d  dy = %d\n",dx,dy)
  picture := make([][]uint8,dy)
  for i := 0;i<len(picture);i++{
    picture[i] = make([]uint8,dx)
  }

   for i := range picture{
     for j:= range picture[i]{
       picture[i][j] = uint8(dy)
       picture[j][i] = uint8(dx)

     }
   }
   return picture 
}

func main(){
  pic.Show(Pic)
}
