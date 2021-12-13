package main
import "fmt"

func main(){
  a,b := 15,20
  b,a = a,b
  fmt.Println(a,b)
}
