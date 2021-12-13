package main

import "fmt"

type Structure struct{
	x int
}

func (structure *Structure) PlusOne(){
	structure.x++
}

func main(){

	structure := Structure{1}
	structure.PlusOne()

	fmt.Println(structure.x)

}
